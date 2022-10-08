
// This is an implementation of Binary Search in java
// PLEASE NOTE THAT BINARY SEARCH WORKS ONLY IN SORTED ARRAYS. IN THE CODE BELOW< WE ASSUME
// IT TO BE SORTED IN ASCENDING ORDER
public class BinarySearch {
    //  Main method
        public static void main(String[] args) {
    //        TEST CASE TO VERIFY CODE
            int[] arr = {1, 4 ,56, 60, 77, 81 , 8686};
              int a = binarySearch(77, arr);
            System.out.println(a);
    //        the output will be 4 (i.e. at index 4)
        }
    
    //     Method for binary search. It returns index if target is found, else returns -1;
         public static  int  binarySearch(int target, int[] arr){
                        int start = 0;
                        int end = arr.length-1;
                        int mid =0;
                        while(start<=end){
                            mid = start+ (end-start)/2;
    
                            if(arr[mid]==target)
                                return mid;
                            else if(arr[mid] <target){
                                start=mid+1;
                            }
                            else{
                                end = mid-1;
                            }
    
                        }
             return -1;
        }
    
    }