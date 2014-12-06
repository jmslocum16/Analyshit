package shit.analyshit;

import android.app.Activity;
import android.content.Context;
import android.os.Bundle;
import android.view.Menu;
import android.view.MenuItem;
import android.widget.ArrayAdapter;
import android.widget.ListView;

import java.util.ArrayList;


public class HomeActivity extends Activity {

    private ListView poopView;
    private CustomAdapter adapter;
    private ArrayList<Poop> poopList;
    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_home);
        poopView = (ListView) findViewById(R.id.poop_view);
        adapter = new CustomAdapter(this, R.layout.list_poop, poopList);
    }


    @Override
    public boolean onCreateOptionsMenu(Menu menu) {
        // Inflate the menu; this adds items to the action bar if it is present.
        getMenuInflater().inflate(R.menu.menu_home, menu);
        return true;
    }

    @Override
    public boolean onOptionsItemSelected(MenuItem item) {
        // Handle action bar item clicks here. The action bar will
        // automatically handle clicks on the Home/Up button, so long
        // as you specify a parent activity in AndroidManifest.xml.
        int id = item.getItemId();

        //noinspection SimplifiableIfStatement
        if (id == R.id.action_settings) {
            return true;
        }

        return super.onOptionsItemSelected(item);
    }

    public class CustomAdapter extends ArrayAdapter<Poop> {
        Context context;
        int resource_id;
        ArrayList<Poop> items;
        public CustomAdapter(Context context, int resource_id, ArrayList<Poop> items) {
            super(context,resource_id,items);
            this.context = context;
            this.resource_id = resource_id;
            this.items = items;
        }
    }
}
