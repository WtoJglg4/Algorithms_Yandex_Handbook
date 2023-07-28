#include <iostream>
#include <algorithm>
using namespace std;

const int len = 200000;

unsigned long long max_product(){
    unsigned long long n, m1 = 1, m2 = 1, cur, ms = 1;
    cin >> n;
    for(int i = 0; i < n; i++){
        cin >> cur;
        if(m1 == 1){
            m1 = cur;
            continue;
        }
        if(m2 == 1){
            m2 = cur;
            ms = m1 * m2;
            if (m2 > m1) swap(m1, m2);
            continue;
        }
        if (cur > m1){
            swap(m1, m2);
            m1 = cur;
            ms = m1 * m2;
            continue;
        }
        if(cur > m2){
            m2 = cur;
            ms = m1 * m2;
        }
    }
    return ms;
}

unsigned long long max_product_sort(){
    unsigned long long n;
    unsigned long long A[len];
    cin >> n;
    for(int i = 0; i < n; i++){
        cin >> A[i];
    }
    sort(A, A + n);
    return A[n-1] * A[n-2];
}

int main()
{
    cout << max_product();
    cout << max_product_sort();
}


