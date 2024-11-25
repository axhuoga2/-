#include <iostream>
#include <cstdlib>
#include <ctime>
#include <string>

std::string generatePassword(int length) {
    const std::string characters = 
        "abcdefghijklmnopqrstuvwxyz"
        "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
        "0123456789"
        "!@#$%^&*()_-+=<>?";
    std::string password;
    int charCount = characters.size();

    // 设置随机数种子
    std::srand(std::time(nullptr));

    for (int i = 0; i < length; ++i) {
        password += characters[std::rand() % charCount];
    }
    return password;
}

int main() {
    int length;

    std::cout << "请输入密码长度：";
    std::cin >> length;

    if (length <= 0) {
        std::cerr << "密码长度必须是正整数！" << std::endl;
        return 1;
    }

    std::string password = generatePassword(length);

    std::cout << "生成的随机密码是：" << password << std::endl;

    return 0;
}
