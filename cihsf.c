#include <pthread.h>
#include <stdio.h>
	// Note the return type: void*

int i = 0;

void* someThreadFunction1(){
	int j = 0;
	for(j; j<1000000;j++){
		i ++;
	}
	return NULL;
}

void* someThreadFunction2(){
	int j = 0;
	for(j; j<1000000;j++){
		i --;
	}
	return NULL;
}


int main(){

	
	pthread_t someThread;
	pthread_t someOtherThread;
	pthread_create(&someThread, NULL, someThreadFunction1, NULL);
	pthread_create(&someOtherThread, NULL, someThreadFunction2, NULL);

	// Arguments to a thread would be passed here ---------^
	pthread_join(someThread, NULL);
	pthread_join(someOtherThread, NULL);

	printf("i = %i", i);

	return 0;
}
