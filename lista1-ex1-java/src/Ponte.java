
import java.util.concurrent.BlockingQueue;
import java.util.concurrent.LinkedBlockingDeque;
import java.util.concurrent.atomic.AtomicBoolean;
import java.util.concurrent.locks.ReentrantLock;

public class Ponte {
    static class Fila {
        private BlockingQueue<Integer> fila;
        String sentido;

        public Fila (String sentido){
            this.fila = new LinkedBlockingDeque<>();
            this.sentido = sentido;
        }

        public void add(Integer value) {
            fila.add(value);
        }

        public Integer dequeue() throws InterruptedException {
            return fila.take();
        }

        public boolean isEmpty() {
            return fila.isEmpty();
        }
    }

    public static void main(String[] args) {
        Fila filaA = new Fila("ino");
        Fila filaB = new Fila("voltano");
        ReentrantLock lock = new ReentrantLock();

        new Thread(() -> {
            try {
                produzirCarros(filaA);
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
        }).start();

        new Thread(() -> {
            try {
                produzirCarros(filaB);
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
        }).start();

        new Thread(() -> {
            try {
                consumirCarros(filaA, lock);
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
        }).start();

        new Thread(() -> {
            try {
                consumirCarros(filaB, lock);
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
        }).start();
    }

    public static void produzirCarros(Fila fila) throws InterruptedException {
        int id = 0;
        while(true) {
            fila.add(id);
            id++;
            Thread.sleep((long)(Math.random() * 100));
        }
    }

    public static void consumirCarros(Fila fila, ReentrantLock ponte) throws InterruptedException {
        while(true) {
            AtomicBoolean hasLock = new AtomicBoolean(false);

            ponte.lock();

            System.out.println("A ponte está liberada para o sentido: " + fila.sentido);

            hasLock.set(true);

            new Thread(() -> {
                try {
                    Thread.sleep(1000);
                    hasLock.set(false);
                } catch (InterruptedException e) {
                    e.printStackTrace();
                }
            }).start();

            while(hasLock.get()) {
                if (!fila.isEmpty()) {
                    Integer carro = fila.dequeue();
                    System.out.println("O carro " + carro + " está " + fila.sentido + " pela ponte");
                    Thread.sleep(200);
                } else {
                    break;
                }
            }

            System.out.println("A ponte está fechada para " + fila.sentido);

            ponte.unlock();
        }
    }
}
