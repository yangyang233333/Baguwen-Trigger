//
// Created by qianyy on 2023/3/13.
//
#include "thread_pool.h"
#include <ctime>
#include <bits/stdc++.h>
#include <unistd.h>
class Blob {
public:
    ~Blob() {
        std::cout << "~Blob" << std::endl;
    }
};

void func() {
    Blob b{};
    std::cout << "------------" << std::endl;
}

int main() {
    std::thread t1(func);

    sleep(10);
}