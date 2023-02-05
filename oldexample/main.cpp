#include "common.h"
#include "Design_Skiplist.h"

int main()
{
    Skiplist kList;
    kList.add(1);
    kList.add(2);
    kList.add(3);

    if (!kList.search(0))
    {
        std::cout << "kList.search(0)" << std::endl;
    }
    kList.add(4);
    if(kList.search(1))   // 返回 true
    {
        std::cout << "kList.search(1)" << std::endl;
    }
    if (!kList.erase(0))   // 返回 false，0 不在跳表中
    {
        std::cout << "kList.search(1)" << std::endl;
    }
    if (kList.erase(1))    // 返回 true
    {
        std::cout << "kList.search(1)" << std::endl;
    }
    if (!kList.search(1))   // 返回 false，1 已被擦除
    {
        std::cout << "kList.search(1)" << std::endl;
    }

    return 0;
}