#include <iostream>
#include <string>

using namespace std;

void byte_order()
{
    int num = 1;

    string order = (1 == *(char *)&num)? "little": "big";
    cout << "Your system use " << order << " endian" << endl;
}

struct ModFlags
{
    char f0: 2;
    char f1: 2;
    char f2: 2;
    char f3: 2;
} mf;

void bit_field_order()
{
    uint8_t *ptr = (uint8_t *)&mf;

    cout << "sizeof(ModFlags): " << sizeof(ModFlags) << endl;
    mf.f0 = 3;
    mf.f3 = 1;
    cout << "mf: 0x" << hex << (int)*ptr << endl;
}

int main() {
    byte_order();
    bit_field_order();

    return 0;
}
