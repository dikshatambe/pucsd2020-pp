#include<stdio.h>
#include<stdlib.h>
#include "../include/calci.h"
int main()
{
int n1,n2;
char operator;
	printf("Enter 2 operands:");
	scanf("%d%d",&n1,&n2);
	printf("Enter operator 1)+ 2)- 3)* 4)/ 5)% 6)q\n");
	
	while(1)
	{
	scanf("%c",&operator);
	
	switch(operator)
	{
		case '1': 
				printf("%d\t",addition(n1,n2));
				break;
		case '2': 
				printf("%d\t",subtraction(n1,n2));
				break;
		case '3': 
				printf("%d\n",multiplication(n1,n2));
				break;
		case '4':
				printf("%f\n",division(n1,n2));
				break;
		case '5':
				if(n1>n2)
                                {
					printf("%f\n",findRemainder(n1,n2));
                                }
				else
				{
					printf("Invalid operands!!");
				}
				break;
		case '6': 
			exit(0);
		
	}
	}

}
