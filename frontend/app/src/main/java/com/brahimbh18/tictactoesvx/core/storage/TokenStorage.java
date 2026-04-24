package com.brahimbh18.tictactoesvx.core.storage;

import android.content.Context;
import android.content.SharedPreferences;

import com.brahimbh18.tictactoesvx.core.util.Constants;

public class TokenStorage {
    private final SharedPreferences prefs;

    public TokenStorage(Context context) {
        this.prefs = context.getSharedPreferences(Constants.SHARED_PREFS, Context.MODE_PRIVATE);
    }

    public void saveToken(String token) {
        prefs.edit().putString(Constants.TOKEN_KEY, token).apply();
    }

    public String getToken() {
        return prefs.getString(Constants.TOKEN_KEY, null);
    }

    public void clear() {
        prefs.edit().remove(Constants.TOKEN_KEY).apply();
    }
}
