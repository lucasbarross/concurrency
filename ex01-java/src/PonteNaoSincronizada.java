
import java.util.concurrent.BlockingQueue;
import java.util.concurrent.LinkedBlockingDeque;
import java.util.concurrent.atomic.AtomicInteger;

public class PonteNaoSincronizada {
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
    }

    public static void main(String[] args) {
        AtomicInteger id = new AtomicInteger(0);
        Fila filaA = new Fila("ino");
        Fila filaB = new Fila("voltano");

        new Thread(() -> {
            try {
                produzirCarros(filaA, id);
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
        }).start();

        new Thread(() -> {
            try {
                produzirCarros(filaB, id);
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
        }).start();

        new Thread(() -> {
            try {
                consumirCarros(filaA);
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
        }).start();

        new Thread(() -> {
            try {
                consumirCarros(filaB);
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
        }).start();
    }

    public static void produzirCarros(Fila fila, AtomicInteger id) throws InterruptedException {
        while(true) {
            int carro = id.addAndGet(1);
            fila.add(carro);
            Thread.sleep((long)(Math.random() * 100));
        }
    }

    public static void consumirCarros(Fila fila) throws InterruptedException {
        while(true) {
            Integer carro = fila.dequeue();
            System.out.println("O carro " + carro + " está " + fila.sentido + " pela ponte");
            Thread.sleep(200);
        }
    }
}
