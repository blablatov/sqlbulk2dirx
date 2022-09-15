using System;
using System.Diagnostics;
using System.ComponentModel;

namespace GoProcessSql
{
    public class MyProcess
    {
		// Library for include to Directum RX, like outside module. Стороняя либра для подключения к Directum RX.
        public static void Main()
        {
            try
            {
                using (Process myProcess = new Process())
                {
                    myProcess.StartInfo.UseShellExecute = false;
                    // Demo directory with the module (sql bulk copy) on Go.
                    // Демо директория размещения исполнительного модуля (sql bulk copy) на Go.
                    myProcess.StartInfo.FileName = "c:\\inetpub\\wwwroot\\insap\\sqlmain.exe";
                    myProcess.StartInfo.CreateNoWindow = true;
                    myProcess.Start();
                    // This code assumes the process you are starting will terminate itself.
                    // Given that it is started without a window so you cannot terminate it
                    // on the desktop, it must terminate itself or you can do it programmatically
                    // from this application using the Kill method.
                }
            }
            catch (Exception e)
            {
                Console.WriteLine(e.Message);
            }
        }
    }
}

