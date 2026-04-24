package com.brahimbh18.tictactoesvx.ui.game;

import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.TextView;

import androidx.annotation.NonNull;
import androidx.recyclerview.widget.RecyclerView;

import com.brahimbh18.tictactoesvx.R;

import java.util.ArrayList;
import java.util.List;

public class BoardAdapter extends RecyclerView.Adapter<BoardAdapter.CellViewHolder> {
    public interface OnCellClickListener { void onCellClick(int index); }

    private final List<String> board = new ArrayList<>();
    private final OnCellClickListener listener;

    public BoardAdapter(OnCellClickListener listener) {
        this.listener = listener;
    }

    public void submit(List<String> cells) {
        board.clear();
        board.addAll(cells);
        notifyDataSetChanged();
    }

    @NonNull
    @Override
    public CellViewHolder onCreateViewHolder(@NonNull ViewGroup parent, int viewType) {
        View view = LayoutInflater.from(parent.getContext()).inflate(R.layout.item_board_cell, parent, false);
        return new CellViewHolder(view);
    }

    @Override
    public void onBindViewHolder(@NonNull CellViewHolder holder, int position) {
        holder.cellText.setText(board.get(position).isEmpty() ? " " : board.get(position));
        holder.itemView.setOnClickListener(v -> listener.onCellClick(position));
    }

    @Override
    public int getItemCount() {
        return board.size();
    }

    static class CellViewHolder extends RecyclerView.ViewHolder {
        final TextView cellText;

        CellViewHolder(@NonNull View itemView) {
            super(itemView);
            cellText = itemView.findViewById(R.id.tvCell);
        }
    }
}
