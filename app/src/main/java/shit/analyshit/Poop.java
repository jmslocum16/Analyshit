package shit.analyshit;

import android.media.Image;

/**
 * Poop Class
 */
public class Poop {
    private Image poop;
    private String date;
    private String log;
    public Poop(Image poop, String date, String log) {
        this.poop = poop;
        this.date = date;
        this.log = log;
    }

    public Image getImage() {
        return poop; }

    public String getDate() {
        return date;
    }

    public String getLog() {
        return log;
    }
}
