import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.util.Scanner;

public class JavaDuplicate {
   public static void main(String args[])
{
  int max=0,n,count=0;
  Scanner s = new Scanner(System.in);
  int no = s.nextInt();
  int [] ar = new int[no];
  for (int k = 0; k < no; k++) 
  {
    ar[k] = s.nextInt();
  }



for(int i=0;i<no;i++)
{
if(ar[i]>max)
{
max=ar[i];
}

}
for(int i=0;i<=max;i++)
{ for(int j=0;j<no;j++)
{
if(i==ar[j])
count++;
}

if(count>1)
System.out.print(i+" ");
count=0;
}
System.out.println();

}
}
