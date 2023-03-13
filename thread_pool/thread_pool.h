//
// Created by qianyy on 2023/3/13.
//
#include <condition_variable>
#include <mutex>
#include <list>
#include <thread>
#include <vector>
#include <iostream>

#ifndef CODE_SEG_THREAD_POOL_H
#define CODE_SEG_THREAD_POOL_H

class Task final {
public:
    void *(*handler)(void *);

    void *args;
};

// 线程池
class ThreadPool final {
private:
    const int max_thread_num_;        // 最大线程数量
    const int max_task_queue_len_;    // 最大任务队列的长度

    std::mutex mu;
    std::condition_variable full_task;
    std::condition_variable empty_task;

    std::list<Task> task_queue;           // 任务队列
    std::vector<std::thread> work_threads; // 工作线程

    bool shutdown;       // 线程池是否被关闭

private:
    void create_thread_pool() {
        for (int i = 0; i < max_thread_num_; ++i) {
            work_threads.emplace_back(&ThreadPool::execute_task, this);
        }
    }

public:
    ThreadPool() = delete;

    ~ThreadPool() {
        destroy();
    }

    // 最大线程数量、最大任务队列长度
    ThreadPool(int max_thread_num, int max_task_queue_len)
            : max_thread_num_(max_thread_num),
              max_task_queue_len_(max_task_queue_len),
              shutdown(false) {
        if (max_thread_num_ > 100) {
            std::cout << "max_thread_num_ > 100 is not allowed." << std::endl;
            return;
        }
        create_thread_pool();
        std::cout << "ThreadPool has created." << std::endl;
    }

    // 添加任务：生产任务
    void add_task(const Task &task) {
        if (shutdown) {
            return;
        }
        std::unique_lock<std::mutex> lock(mu);
        // 如果任务队列已满，则挂起当前线程，等待不满的时候被唤醒
        while (task_queue.size() == max_task_queue_len_ && !shutdown) {
            full_task.wait(lock); // 等待任务被消耗
            if (shutdown) { // 唤醒后检查是否被关闭
                return;
            }
        }
        task_queue.push_back(task);
        empty_task.notify_one(); // 生产了一个task，唤醒因空task队列而挂起的线程
    }

    // 执行任务：消费任务
    void execute_task() {
        while (!shutdown) {
            std::unique_lock<std::mutex> lock(mu);
            while (task_queue.empty()) {
                empty_task.wait(lock);
                if (shutdown) {
                    return;
                }
            }

            auto task = task_queue.front();
            task_queue.pop_front();
            full_task.notify_one(); // 消费了一个线程后，唤醒因task_queue满而等待的线程
            lock.unlock(); // 获取到任务后提前解锁，降低锁范围

            // 执行任务
            (task.handler)(task.args);
        }
    }

    // 销毁线程池
    void destroy() {
        if (shutdown) {
            return;
        }
        shutdown = true;

        // 唤醒所有线程池
        empty_task.notify_all();
        full_task.notify_all();

        for (int i = 0; i < max_thread_num_; ++i) {
            work_threads[i].join();
        }

        task_queue.clear();
        work_threads.clear();
    }
};

#endif //CODE_SEG_THREAD_POOL_H
