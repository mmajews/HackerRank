import java.io.*;
import java.util.*;
import java.text.*;
import java.math.*;
import java.util.regex.*;

public class Solution {

    static void miniMaxSum(int[] arr) {
        List<Integer> elements = new ArrayList<>();
        for (int i : arr) {
            elements.add(i);
        }

        Collections.sort(elements);

        long minimalSum = 0;
        long maximumSum = 0;


        for (int i = 0; i < 4; i++) {
            minimalSum += elements.get(i);
        }

        for (int i = elements.size() - 1; i > elements.size() - 5; i--) {
            maximumSum += elements.get(i);
        }

        System.out.print(minimalSum);
        System.out.print(" ");
        System.out.print(maximumSum);
    }

    public static void main(String[] args) {
        Scanner in = new Scanner(System.in);
        int[] arr = new int[5];
        for (int arr_i = 0; arr_i < 5; arr_i++) {
            arr[arr_i] = in.nextInt();
        }
        miniMaxSum(arr);
        in.close();
    }
}
