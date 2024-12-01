import matplotlib.pyplot as plt
import matplotlib.animation as animation
import numpy as np
import seaborn as sns

# open csv
with open('input.dat') as file:
    info = file.readline().rstrip('\n').split()

    start = float(info[0])
    stop = float(info[1])
    delta = float(info[2])
    h = int(info[3])
    w = int(info[4])

    print(f" start: {start}, stop: {stop}, delta: {delta}, h: {h}, w: {w}")

    frames = int((stop-start)/delta)
    lines = file.readlines()
    data = np.zeros((int((stop-start)/delta)+1, h, w))
    for t in range(int(start), int(stop)+1, int(delta)):
        for i in range(h):
            for j in range(w):
                time, value = lines[t*h*w + i*h + j].rstrip('\n').split()
                data[t][i][j] = float(value)


def gen_frame(i):
    if i < start:
        print(f"frame out of bounds -- showing zeros. frame: {i}")
        return data[0]

    if i > stop:
        print(f"frame out of bounds -- showing zeros. frame: {i}")
        return data[-1]

    return data[i]


grid_kws = {'width_ratios': (0.9, 0.05), 'wspace': 0.2}
fig, (ax, cbar_ax) = plt.subplots(1, 2, gridspec_kw=grid_kws, figsize=(12, 12))


def update(frame):
    ax.clear()
    frame_data = gen_frame(frame)
    print(f"frame: {frame}")
    sns.heatmap(frame_data,
                ax=ax,
                cbar_ax=cbar_ax,
                cbar=True,
                vmin=data.min(),
                vmax=data.max())
    ax.set_title(f"frame {frame}")


anim = animation.FuncAnimation(fig, update, frames=frames)
anim.save('./heatmap.mp4', writer='ffmpeg', fps=10)
