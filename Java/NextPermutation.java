class Solution {
    public void nextPermutation(int[] nums) {
        int n =  nums.length;
        int i = n-2;
        while(i >= 0 && nums[i] >= nums[i+1]){
            i--;
        }
        if(i >= 0){
            int j = n-1;
        
            while(j >= 0 && nums[i] >= nums[j])
                j--;

            int t = nums[i];
            nums[i]  = nums[j];
            nums[j] = t;   
        }
        reverse(nums, i+1);
    }
    
    void reverse(int[] nums, int i){
        int j = nums.length-1;
        while(i < j){
            int t = nums[i];
            nums[i] = nums[j];
            nums[j] = t;
            i++;
            j--;
        }
    }
}