using UnityEngine;
using System.Collections;
using System.Net;
using System.Net.Sockets;

public class netdemo : MonoBehaviour {
    TcpClient client;
    NetworkStream stream;
    byte[] bts = new byte[2048];
    string laststr="";

    void Start () {
        client = new TcpClient();
        
    }

    void clearBts()
    {
        for (int i = 0; i < bts.Length; i++) {
            bts[i] = 0;
        }
    }
		
	void Update () {
        if (client.Available > 0) {
            clearBts();
            stream.Read(bts, 0, bts.Length);
            laststr = System.Text.Encoding.ASCII.GetString(bts);
        }
	}

    string sendstr="";
    void OnGUI()
    {
        if (!client.Connected)
        {
            if (GUI.Button(new Rect(200, 200, 200, 50), "连接"))
            {
                client.Connect("127.0.0.1", 8008);
            }
            return;
        }

        GUI.Label(new Rect(200, 0, 200, 50), laststr);
        if(stream == null)
            stream = client.GetStream();
        sendstr = GUI.TextField(new Rect(200, 100, 200, 50), sendstr);
        if (string.IsNullOrEmpty(sendstr)) return;
        if (GUI.Button(new Rect(200, 200, 200, 50), "发送"))
        {
            byte[] data = System.Text.Encoding.ASCII.GetBytes(sendstr);
            stream.Write(data, 0, data.Length);
        }
        GUI.Label(new Rect(200, 300, 200, 50), "len=" + client.Available.ToString());       
    }

    void OnDestroy()
    {
        client.Close();
    }
}
