#!/usr/bin/env python3
import os
import sys
import time
import json
import subprocess
import unicodedata
from datetime import datetime

# 端末サイズ取得
try:
    screen_cols = os.get_terminal_size().columns
except OSError:
    screen_cols = 80  # デフォルト幅

# 全角/半角文字の幅を計算
def char_width(ch):
    return 2 if unicodedata.east_asian_width(ch) in ("W", "F") else 1

# 表示幅ベースで substring を取得
def substr_by_width(s, offset, width):
    disp_index = 0
    out = ""
    for ch in s:
        w = char_width(ch)
        if disp_index + w <= offset:
            disp_index += w
            continue
        if disp_index >= offset + width:
            break
        out += ch
        disp_index += w
    return out

# 文字列の表示幅を計算
def display_width(s):
    return sum(char_width(ch) for ch in s)

# ニュース取得
def fetch_news():
    try:
        cmd = [
            "curl", "-s",
            "https://api.rss2json.com/v1/api.json?rss_url=https://www3.nhk.or.jp/rss/news/cat0.xml"
        ]
        result = subprocess.run(cmd, capture_output=True, text=True, timeout=10)

        if result.returncode == 0:
            data = json.loads(result.stdout)
            if data.get("status") == "ok":
                items = data.get("items", [])
                news_list = []
                for item in items[:8]:  # 最新8件
                    title = item.get("title", "")
                    if title:
                        news_list.append(f"■ {title}")
                return news_list
    except (subprocess.TimeoutExpired, json.JSONDecodeError, Exception):
        pass

    return ["■ ニュースの取得に失敗しました", "■ インターネット接続を確認してください"]

# 読み上げ機能
def speak_text(text):
    try:
        clean_text = text.replace("■ ", "")
        subprocess.Popen(["say", clean_text, "-v", "Kyoko", "-r", "160"])
    except Exception:
        pass  # 読み上げ失敗時は無視

# 1つのニュースをスクロール表示
def scroll_news_item(text):
    # 読み上げ開始
    speak_text(text)

    # スクロール用文字列作成
    spaces = " " * screen_cols
    full_display = spaces + text + spaces
    total_width = display_width(full_display)

    # スクロール実行
    for i in range(total_width - screen_cols + 1):
        # 6行目に移動してクリア
        print("\033[6;0H\033[K", end="")
        segment = substr_by_width(full_display, i, screen_cols)
        print(f"\033[1;33m{segment}\033[0m", end="")

        # 7行目もクリア（安全対策）
        print("\033[7;0H\033[K", end="")

        sys.stdout.flush()
        time.sleep(0.04)

    # 完了後少し待機
    time.sleep(0.5)

# ヘッダー生成
def generate_header():
    print("\033[2J\033[H", end="")  # 画面クリア
    line = "━" * screen_cols
    print(f"\033[1;34m{line}\033[0m")

    title = "📺🔊 ライブニュースティッカー"
    header = title.center(screen_cols)
    print(f"\033[1;36m{header}\033[0m")
    print(f"\033[1;34m{line}\033[0m")
    print()
    print("🔄 永続的にニュースを取得・表示・読み上げします")
    print("⏱️  5分間隔で新しいニュースを取得")
    print("🎯 Ctrl+C で終了")
    print()

# メイン処理
def main():
    generate_header()

    # カーソル非表示
    print("\033[?25l", end="")

    news_cache = []
    current_index = 0
    last_fetch_time = 0

    try:
        while True:
            current_time = time.time()

            # 5分間隔でニュース更新（または初回）
            if current_time - last_fetch_time > 300 or not news_cache:
                print("\033[4;0H\033[K🔄 ニュースを更新中...", end="")
                sys.stdout.flush()

                news_cache = fetch_news()
                current_index = 0
                last_fetch_time = current_time

                print(f"\033[4;0H\033[K✅ {len(news_cache)}件のニュースを取得")
                time.sleep(1)

            # ニュース表示
            if news_cache:
                current_news = news_cache[current_index]

                # ニュース情報表示
                now = datetime.now().strftime("%H:%M")
                info = f"🔄 ニュース {current_index + 1}/{len(news_cache)} | 最終更新: {now}"
                print(f"\033[4;0H\033[K\033[1;32m{info}\033[0m")

                # スクロール&読み上げ
                scroll_news_item(current_news)

                # 次のニュースへ
                current_index = (current_index + 1) % len(news_cache)

                # 全ニュース完了後は少し休憩
                if current_index == 0:
                    time.sleep(3)
            else:
                time.sleep(5)

    except KeyboardInterrupt:
        print("\033[?25h")  # カーソル復元
        print("\n📺 ニュースティッカーを終了しました")
        # 読み上げプロセスを停止
        try:
            subprocess.run(["killall", "say"], capture_output=True)
        except:
            pass

if __name__ == "__main__":
    main()