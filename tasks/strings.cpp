#include <iostream>
using namespace std;

const int len = 20;

void strings(int& n, char* A){
    cin >> n;
    for(int i=0; i<2*n; i= i+ 2){
        cin >> A[i];
    }
    for(int i=1; i<2*n; i= i+ 2){
        cin >> A[i];
    }
}

int main(){
    int n;
    char A[len];
    strings(n, A);

    for(int i=0; i < 2*n; i++){
        cout << A[i];
    }
}