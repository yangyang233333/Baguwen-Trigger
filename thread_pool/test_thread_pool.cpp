//
// Created by qianyy on 2023/3/13.
//
#include "thread_pool.h"
#include <ctime>
#include <bits/stdc++.h>
#include <unistd.h>

void *handler(void *args) {
    int a = *(int *) args;
    std::cout << "a*a=" << a * a << std::endl;
    return nullptr;
}

int main() {
    ThreadPool pools(5, 10);
    std::vector<int> datas = {1, 2, 3, 4, 5, 6, 7, 8, 9};
    for (int &data: datas) {
        pools.add_task(Task{handler, &data});
    }

    sleep(3);
    std::reverse(datas.begin(), datas.end());
    for (int &data: datas) {
        pools.add_task(Task{handler, &data});
    }
    sleep(3);
    pools.destroy();
}