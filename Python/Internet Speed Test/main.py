import speedtest
import sys

def speedtester():
    try:
        print("\n🔍 Initializing speed test...\n")
        st = speedtest.Speedtest()

        print("📡 Loading server list...")
        st.get_servers()
        print("🌍 Selecting best server...")
        best = st.get_best_server()
        print(f"Connected to: {best['host']} ({best['country']})\n")

        while True:
            try:
                print("Choose a test to perform:")
                print("1️⃣  Download Speed")
                print("2️⃣  Upload Speed")
                print("3️⃣  Ping")
                print("4️⃣  Run Full Test")
                print("0️⃣  Exit")

                option = int(input("\nEnter your choice: "))

                if option == 0:
                    print("Exiting... Goodbye!")
                    sys.exit(0)

                elif option == 1:
                    print("\nTesting download speed...")
                    download_speed = st.download() / 1_000_000
                    print(f"Download Speed: {download_speed:.2f} Mbps\n")

                elif option == 2:
                    print("\nTesting upload speed...")
                    upload_speed = st.upload() / 1_000_000
                    print(f"Upload Speed: {upload_speed:.2f} Mbps\n")

                elif option == 3:
                    ping_result = st.results.ping
                    print(f"\nPing: {ping_result:.2f} ms\n")

                elif option == 4:
                    print("\n Running full test...")
                    download_speed = st.download() / 1_000_000
                    upload_speed = st.upload() / 1_000_000
                    ping_result = st.results.ping

                    print("\n===== SPEED TEST RESULTS =====")
                    print(f"Download Speed: {download_speed:.2f} Mbps")
                    print(f"Upload Speed  : {upload_speed:.2f} Mbps")
                    print(f"Ping          : {ping_result:.2f} ms")
                    print("==============================\n")

                else:
                    print(" Invalid choice! Please select a valid option.\n")

            except ValueError:
                print("Please enter a number (0–4).\n")

    except speedtest.ConfigRetrievalError:
        print("Failed to retrieve configuration. Please check your connection.")
    except Exception as e:
        print(f"An unexpected error occurred: {e}")


if __name__ == "__main__":
    speedtester()
