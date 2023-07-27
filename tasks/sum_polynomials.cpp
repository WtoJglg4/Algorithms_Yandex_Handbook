#include <iostream>
using namespace std;

const int len = 10;

//int Sum_polynomials(){
int* Sum_polynomials(int& n, int& m, int* A, int* B){
  
    cin>>n;
    for(int i=0; i<=n; i++){
        cin >> A[i];
    }
    cin>>m;
    for(int i=0; i<=m; i++){
        cin >> B[i];
    }
    int k = abs(n-m);
    //cout << "\nk: " << k << endl;
    if (n >= m){
        for(int i=k; i<=n + k; i++){
            A[i] += B[i-k];
            //cout << A[i] << " ";
        }
        return A;
    }
    else{
        for(int i=k; i<=m + k; i++){
            B[i] += A[i-k];
            //cout << B[i] << " ";
        }
        return B;
    }
}

int main(){
    int n = 0, m = 0;
    int A[len] = {101};
    int B[len] = {101};

    int* arr = Sum_polynomials(n, m, A, B);
    //cout << arr << endl;
    int l = max(n, m);

    cout << l << endl;
    for(int i=0; i<=l; i++){
        cout << arr[i] << " ";
    }
}