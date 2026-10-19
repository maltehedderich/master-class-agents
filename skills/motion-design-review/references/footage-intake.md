# Footage intake

Use this reference when the review material is a video file. It turns the clip into frames and measurements you can cite by timestamp. Write outputs to a scratch or temporary directory, not the user's project, unless they ask otherwise.

The commands need `ffmpeg` and `ffprobe`. If they are missing, ask before installing (`brew install ffmpeg` on macOS), or ask the user for exported frames at 2 fps, the first and last frames, and the cut timestamps.

```sh
IN="path/to/clip.mp4"
OUT="path/to/scratch/motion-review"
mkdir -p "$OUT"
```

## 1. Probe the file

```sh
ffprobe -v error -show_entries format=duration:stream=codec_type,width,height,r_frame_rate:stream_side_data=rotation -of compact=p=0 "$IN"
```

Record duration, resolution, aspect ratio, frame rate, and whether an audio stream exists. A `rotation=90` or `rotation=-90` value means the displayed orientation swaps width and height. No audio stream means the clip is silent by design, so sound points under R7 do not apply.

## 2. Contact sheet

```sh
ffmpeg -v error -i "$IN" -vf "fps=1,scale=270:-2,tile=5x4:padding=4:color=white" "$OUT/sheet_%02d.png"
```

Read tiles left to right, top to bottom. Tile `n` (counting from 0) on sheet `k` (counting from 1) shows `t ≈ ((k − 1) × 20 + n) / fps` seconds, which at `fps=1` is simply second `n` on the first sheet. Use `fps=2` for clips under about 12 seconds.

Burned-in timestamps need the `drawtext` filter, which is missing from many ffmpeg builds. Work out the time from tile position instead of depending on it.

The sheet is for structure and gaze path. Do not judge legibility from it; tiles are far smaller than viewing size.

## 3. Frames for close reading

```sh
ffmpeg -v error -i "$IN" -vf "fps=2,scale=-2:960" "$OUT/f_%03d.png"
```

Frame `f_NNN.png` shows `t ≈ (NNN − 1) / 2` seconds. Use these to cite timestamps and to catch text that appears only briefly. Raise `fps` to 4–5 to step through a fast sequence.

## 4. First, last, and shot frames

```sh
ffmpeg -v error -i "$IN" -frames:v 1 "$OUT/first.png"
ffmpeg -v error -sseof -1 -i "$IN" -update 1 "$OUT/last.png"
ffmpeg -v error -i "$IN" -vf "select='eq(n,0)+gt(scene,0.3)',scale=-2:960" -fps_mode vfr "$OUT/cut_%02d.png"
```

`first.png` is the frozen-first-frame test for R2. `last.png` is the ending for R8 and one half of the loop-seam check. `cut_NN.png` is the first frame of each shot, in the same order as the cut list below. On ffmpeg older than 5.1, replace `-fps_mode vfr` with `-vsync vfr`.

To grab the frame at a specific moment, such as the middle of a hold:

```sh
ffmpeg -v error -ss 8.5 -i "$IN" -frames:v 1 "$OUT/at_8.5.png"
```

## 5. Cut list

```sh
ffmpeg -hide_banner -i "$IN" -vf "select='gt(scene,0.3)',metadata=print:file=-" -an -f null - 2>/dev/null | grep -o 'pts_time:[0-9.]*'
```

Each `pts_time` is a hard cut or a large visual change. Shot lengths are the gaps between consecutive cuts, with the clip start and end as the outer bounds. Lower the threshold to `0.2` for subtle edits; raise it to `0.4`–`0.5` if UI animation produces false cuts. Continuous camera moves and morph transitions usually do not register.

## 6. Holds

```sh
ffmpeg -hide_banner -i "$IN" -vf "freezedetect=n=-60dB:d=0.5" -map 0:v:0 -f null - 2>&1 | grep -o 'freeze_[a-z]*: [0-9.]*'
```

This reports stretches of at least 0.5 s where the frame barely changes. A `freeze_start` with no matching `freeze_end` runs to the end of the clip. Use these to check whether results and key text stay on screen long enough (R6).

The measure is whole-frame change, so a small moving element such as a cursor can sit inside a reported hold. Check the frames in that range before calling it a rest. To separate fully still frames from near-holds, rerun with `n=-80dB`. If the stricter run finds nothing in real footage, compression noise is the likely cause; rely on the `-60dB` result and frame inspection.

## 7. Audio

```sh
ffmpeg -hide_banner -i "$IN" -af "silencedetect=noise=-40dB:d=0.5" -vn -f null - 2>&1 | grep -o 'silence_[a-z]*: [0-9.]*'
```

This shows only where the track is silent. You cannot hear the audio. For VO, music, and sound-effect judgments, ask for the script, a transcript, or sound notes, and align VO lines to the beat map by timestamp. Without them, mark the sound parts of R7 `n/a`.

## 8. Viewing size, loop seam, and easing

```sh
# Approximate a phone feed (about 390 px wide)
ffmpeg -v error -i "$OUT/cut_02.png" -vf "scale=390:-2" "$OUT/cut_02_phone.png"

# Loop seam: first and last frames side by side
ffmpeg -v error -i "$OUT/first.png" -i "$OUT/last.png" -filter_complex "[0][1]hstack,scale=-2:960" "$OUT/loop_seam.png"

# Every frame of a 0.6 s transition starting at 4.0 s (18 frames at 30 fps)
ffmpeg -v error -ss 4.0 -t 0.6 -i "$IN" -vf "scale=-2:480,tile=6x3:padding=4:color=white" -frames:v 1 "$OUT/transition_4.0.png"
```

- **Viewing size:** scale key frames to the real display width (about 390 px for a phone feed, or the embed's rendered width on a landing page). If essential text is hard to read in the scaled frame, it is hard to read in the feed.
- **Loop seam:** for landing-page loops, the two halves should match closely enough that the restart is invisible.
- **Easing:** compare an element's position from frame to frame across the transition strip. Even steps mean linear motion; steps that shrink toward the arrival mean ease-out; jumps across most of the frame in one or two steps mean the move is too fast to follow. Size the tile grid to the frame count (`-t` × frame rate).

## Report what you measured

Cite measured times in findings, for example "the result holds 0.6 s (9.2–9.8 s)". State the limits: sampled frames miss what happens between samples, the scene threshold is approximate, and easing read from a transition strip is a rough approximation of seeing the clip play. When a judgment depends on playback you could not do, mark it `n/a` or ask the user to describe it.
