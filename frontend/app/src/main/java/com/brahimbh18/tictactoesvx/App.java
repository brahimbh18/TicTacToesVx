package com.brahimbh18.tictactoesvx;

import android.app.Application;

import com.brahimbh18.tictactoesvx.core.di.ServiceLocator;

public class App extends Application {
    @Override
    public void onCreate() {
        super.onCreate();
        ServiceLocator.init(this);
    }
}
