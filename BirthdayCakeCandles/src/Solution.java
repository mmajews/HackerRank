import java.io.*;
import java.util.*;
import java.text.*;
import java.math.*;
import java.util.regex.*;

public class Solution {

    static int birthdayCakeCandles(int n, int[] ar) {
        OptionalInt max = Arrays.stream(ar).max();
        if(!max.isPresent()){
            throw new RuntimeException("Array cannot be empty!");
        }
        int maximum = max.getAsInt();

        //Can be safely casted to int, requirements of the task
        return (int) Arrays.stream(ar).filter(candleHeight -> candleHeight == maximum).count();
    }

    public static void main(String[] args) {
        Scanner in = new Scanner(System.in);
        int n = in.nextInt();
        int[] ar = new int[n];
        for(int ar_i = 0; ar_i < n; ar_i++){
            ar[ar_i] = in.nextInt();
        }
        int result = birthdayCakeCandles(n, ar);
        System.out.println(result);
    }
}
