#include <iostream>

using namespace std;

void basic_type_size()
{
    cout << (8*sizeof(void *)) << " bit system" << endl;
    cout << "char: " << sizeof(char) << endl;
    cout << "short: " << sizeof(short) << endl;
    cout << "int: " << sizeof(int) << endl;
    cout << "long: " << sizeof(long) << endl;
    cout << "long long: " << sizeof(long long) << endl;
    cout << "pointer: " << sizeof(void *) << endl;
    cout << "float: " << sizeof(float) << endl;
    cout << "double: " << sizeof(double ) << endl;
}

struct DefSt{
    char a;
    int b;
    short c;
}ds[2];
void structSizeDef()
{
    cout << "sizeof(DefSt): " << sizeof(DefSt) << endl;
    cout << "ds[0] @" << (void *)&ds[0] << endl;
    cout << "ds[0].a @" << (void *)&ds[0].a << endl;
    cout << "ds[0].b @" << (void *)&ds[0].b << endl;
    cout << "ds[0].c @" << (void *)&ds[0].c << endl;
    cout << "ds[1] @" << (void *)&ds[1] << endl;
}

#pragma pack(1)
struct SpecSt{
    char a;
    int b;
    short c;
} ss0, ss1;
#pragma pack()
void structSizeSpec()
{
    cout << "sizeof(SpecSt): " << sizeof(SpecSt) << endl;
    cout << "ss0 @" << (void *)&ss0 << endl;
    cout << "ss0.a @" << (void *)&ss0.a << endl;
    cout << "ss0.b @" << (void *)&ss0.b << endl;
    cout << "ss0.c @" << (void *)&ss0.c << endl;
    cout << "ss1 @" << (void *)&ss1 << endl;
}

int main()
{
    basic_type_size();
    structSizeDef();
    structSizeSpec();

    return 0;
}
