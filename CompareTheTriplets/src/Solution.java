import java.util.Scanner;

public class Solution {

    static int[] solve(int a0, int a1, int a2, int b0, int b1, int b2) {
        // Complete this function
        int aPoints = 0;
        int bPoints = 0;

        int[] allNumbersA = new int[]{a0, a1, a2};
        int[] allNumbersB = new int[]{b0, b1, b2};
        for(int i = 0 ; i< allNumbersA.length; i++){
            if (allNumbersA[i] > allNumbersB[i]) {
                aPoints += 1;
            } else if (allNumbersB[i] > allNumbersA[i]) {
                bPoints += 1;
            }
        }

        return new int[]{aPoints, bPoints};
    }

    public static void main(String[] args) {
        Scanner in = new Scanner(System.in);
        int a0 = in.nextInt();
        int a1 = in.nextInt();
        int a2 = in.nextInt();
        int b0 = in.nextInt();
        int b1 = in.nextInt();
        int b2 = in.nextInt();
        int[] result = solve(a0, a1, a2, b0, b1, b2);
        for (int i = 0; i < result.length; i++) {
            System.out.print(result[i] + (i != result.length - 1 ? " " : ""));
        }
        System.out.println("");
    }
}