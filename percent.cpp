#include <iostream>
using namespace std;
#include <stdlib.h>
#include <time.h>
double per[26] = {8.167,1.492,2.782,4.253,12.702,2.228,2.015,6.094,6.966,0.153,0.772,4.025,2.406,6.749,7.507,1.929,0.095,5.987,6.327,9.056,2.758,0.978,2.360,0.150,1.974,0.074};
char let[26] = {'a','b','c','d','e','f','g','h','i','j','k','l','m','n','o','p','q','r','s','t','u','v','w','x','y','z'};
long counts[26] = {};
const long MAX = 100001L;
void init(){
    srand((unsigned)time(NULL));
    long sum = 0L;
    fill(counts,counts+26,0L);
    for(int i=0;i<26;i++){
        sum += long(per[i]*1000L);
        counts[i] = sum;
    }
    counts[26]=MAX;
}
char* output(int length){
    char *ret = new char(length+1);
    for(int i=0;i<=length;i++)
        ret[i]='\0';
    for(int i=0;i<length;i++){
        long v = long((double)rand()/RAND_MAX*MAX);
        v %= MAX;
        //cout<<v<<endl;
        int k = 0;
        for(k=0;k<26;k++){
            if(counts[k]>=v)
                break;
        }
        k %= 26;
        //cout<<let[k]<<"\t";
        ret[i] = let[k];
    }
    //cout<<endl;
    return ret;
}
void display(){
    for(int i=0;i<26;i++){
        cout<<let[i]<<":"<<counts[i]<<endl;
    }
    cout<<endl;
}
void _test_generate_word(const int&word_num, const int&max_word_length){
	unsigned int count[26];
	double sta_pro[26];
	unsigned long long word_count = 0;
	fill(count, count + 26, 0);
	fill(sta_pro, sta_pro + 26, 0);

	int i = 0;
	char *word = NULL;
	while (i < word_num){
		word = output(rand() % (max_word_length)+1);
		for (int j = 0; word[j] != '\0'; j++){
			count[word[j] - 'a']++;
			word_count++;
		}
		i++;
	}
	for (int i = 0; i < 26; i++){
		sta_pro[i] = (double)count[i] / word_count;
		cout << (char)('a' + i) << ": " << sta_pro[i]*100 << endl;
	}
}
int main()
{
    init();
    display();
    _test_generate_word(10000,10);
    return 0;
}
