#!/usr/bin/env python3
import os
import sys
import time
import json
import subprocess
import unicodedata
import random
from datetime import datetime

# ニュースソース設定
NEWS_SOURCES = [
    {
        "name": "NHK",
        "rss_url": "https://www3.nhk.or.jp/rss/news/cat0.xml",
        "max_items": 8
    },
    {
        "name": "お笑いナタリー",
        "rss_url": "https://natalie.mu/owarai/feed/news",
        "max_items": 5
    },
    {
        "name": "時事通信",
        "rss_url": "https://www.jiji.com/rss/ranking.rdf",
        "max_items": 5
    }
]

# 表示モード設定
DISPLAY_MODE_RANDOM = True  # True: ランダム表示, False: 順次表示

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

# 単一ソースからニュース取得
def fetch_news_from_source(source):
    try:
        cmd = [
            "curl", "-s",
            f"https://api.rss2json.com/v1/api.json?rss_url={source['rss_url']}"
        ]
        result = subprocess.run(cmd, capture_output=True, text=True, timeout=10)

        if result.returncode == 0:
            data = json.loads(result.stdout)
            if data.get("status") == "ok":
                items = data.get("items", [])
                news_list = []
                for item in items[:source['max_items']]:
                    title = item.get("title", "")
                    if title:
                        # ソース名を末尾に追加
                        formatted_news = f"■ {title} ({source['name']})"
                        news_list.append(formatted_news)
                return news_list
    except (subprocess.TimeoutExpired, json.JSONDecodeError, Exception):
        pass

    return []

# 全ソースからニュース取得
def fetch_all_news():
    all_news = []
    failed_sources = []

    for source in NEWS_SOURCES:
        print(f"\033[4;0H\033[K🔄 {source['name']}からニュースを取得中...", end="")
        sys.stdout.flush()

        news_from_source = fetch_news_from_source(source)
        if news_from_source:
            all_news.extend(news_from_source)
        else:
            failed_sources.append(source['name'])

        time.sleep(0.5)  # ソース間で少し待機

    # 取得に失敗したソースがある場合の通知
    if failed_sources:
        failed_msg = f"■ 取得失敗: {', '.join(failed_sources)}"
        all_news.append(failed_msg)

    # 何も取得できなかった場合
    if not all_news:
        return ["■ ニュースの取得に失敗しました", "■ インターネット接続を確認してください"]

    # 表示モードに応じてソート
    if DISPLAY_MODE_RANDOM:
        # ランダム表示（取得失敗メッセージは最後に固定）
        error_messages = [msg for msg in all_news if "取得失敗" in msg]
        news_messages = [msg for msg in all_news if "取得失敗" not in msg]
        random.shuffle(news_messages)
        all_news = news_messages + error_messages

    return all_news

# 読み上げ機能
def speak_text(text):
    try:
        # ソース名部分を除去してから読み上げ
        clean_text = text.replace("■ ", "")
        # (ソース名) の部分を除去
        import re
        clean_text = re.sub(r'\s*\([^)]+\)\s*$', '', clean_text)
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
    print(f"📺 対応ソース: {', '.join([source['name'] for source in NEWS_SOURCES])}")
    print(f"🎲 表示モード: {'ランダム' if DISPLAY_MODE_RANDOM else '順次'}")
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

                news_cache = fetch_all_news()
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