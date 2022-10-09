class Solution {
    public String longestPalindrome(String s) {
        int n = s.length();
        boolean[][] dp = new boolean[n][n];
        
        for(int i = 0; i < n; i++){
            dp[i][i] = true;
        }
        int max = 0, ind = 0;
        for(int i = 1; i < n; i++){
            for(int j = 0; j < n - i; j++){
                if(s.charAt(j) == s.charAt(i+j)){
                    if(i != 1){
                        dp[j][i+j] = dp[j+1][i+j-1];
                    }
                    else
                        dp[j][i+j] = true;
                }
                if(dp[j][i+j] && i > max){
                    max = i;
                    ind = j;
                }
            }
        }
        return s.substring(ind, ind+max+1);
    }
}