import java.util.concurrent.*;

import static java.lang.System.out;

public class LegApp {
    public static void main(String[] args) throws InterruptedException {
        var semaLeft = new Semaphore(1);
        var semaRight = new Semaphore(1);
        semaRight.acquire();
        var left = new Thread(new Leg("left", semaLeft, semaRight));
        left.setDaemon(true);
        left.start();
        var right = new Thread(new Leg("right", semaRight, semaLeft));
        right.setDaemon(true);
        right.start();
        Thread.sleep(1000);
    }

    static class Leg implements Runnable {
        final String kind;
        final Semaphore sema, next;

        public Leg(String kind, Semaphore sema, Semaphore next) {
            this.kind = kind;
            this.sema = sema;
            this.next = next;
        }

        @Override
        public void run() {
            while (true) {
                try {
                    sema.acquire();
                } catch (InterruptedException e) {
                    break;
                }
                out.println(kind + " step");
                next.release();
            }
        }
    }
}

