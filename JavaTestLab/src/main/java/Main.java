public class Main {
    public static void main(String[] args) {
        for (int i = 0; i < 10; i++) {
            clearConsole();
            System.out.println("Hello world!" + i);
        }
    }

    public final static void clearConsole() {
        try {
            final String os = System.getProperty("os.name").toLowerCase();

            if (os.contains("win")) {
                Runtime.getRuntime().exec("cls");
            } else if (os.contains("mac")) {
                Runtime.getRuntime().exec("clear");
            } else {
                System.out.print("\033\143");
            }
        } catch (Exception e) {
            e.printStackTrace();
        }
    }
}