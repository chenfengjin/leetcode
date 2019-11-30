/*
 * @lc app=leetcode id=66 lang=c
 *
 * [66] Plus One
 */


/**
 * Note: The returned array must be malloced, assume caller calls free().
 */
int* plusOne(int* digits, int digitsSize, int* returnSize){
    int carry = 1;
    for(int i = digitsSize-1; i >= 0;i--){
        digits[i]=(digits[i]+carry)%10;
        carry=(digits[i]+carry)/10;
    }
    *returnSize = digitsSize;

    int* returnDigits;
    // if (carry){
    //     *returnSize++;
    //     returnDigits = malloc(sizeof(int)*(*returnSize));
    //     returnDigits[0]=1;          
    // }else{
        returnDigits = malloc(sizeof(int)*(*returnSize));
    // }
    for(int i = 0 ;i<digitsSize;i++){
        returnDigits[i+1]=digits[i];
    } 
    return returnDigits;
}



