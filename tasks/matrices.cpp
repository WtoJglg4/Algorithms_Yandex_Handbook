#include <iostream>
using namespace std;

const int len = 10;

void matrices(int& n, int& m, int A[len][len]){
    cin >> n >> m;
    for(int k=0; k < 2; k++){
        for(int i=0; i < n; i++){
            for(int j=0; j < m; j++){
                int temp;
                cin >> temp;
                A[i][j] += temp;
            }
        }
    }  
}

int main()
{
    int n, m;
    int A[len][len] = {0};
    matrices(n, m, A);

    for(int i=0; i < n; i++){
        for(int j=0; j < m; j++){
            cout << A[i][j] << " ";
        }
        cout << endl;
    }
}


