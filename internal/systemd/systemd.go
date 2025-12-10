package systemd

import (
    "os"
    "path/filepath"
    "rustdesk_install/internal/util"
)

func HasSystemctl() bool {
    return util.HasBinary("systemctl")
}

func SignalUnit(installDir, logDir, hbsPort, idAddr, hbrPort string) string {
    return "[Unit]\n" +
        "Description=RustDesk Signal Server\n" +
        "After=network.target\n" +
        "[Service]\n" +
        "Type=simple\n" +
        "ExecStart=" + filepath.Join(installDir, "hbbs") + " -p " + hbsPort + " -r " + idAddr + ":" + hbrPort + "\n" +
        "WorkingDirectory=" + installDir + "\n" +
        "Restart=always\n" +
        "RestartSec=10\n" +
        "StandardOutput=append:" + filepath.Join(logDir, "signal.log") + "\n" +
        "StandardError=append:" + filepath.Join(logDir, "signal.error") + "\n" +
        "[Install]\n" +
        "WantedBy=multi-user.target\n"
}

func RelayUnit(installDir, logDir, hbrPort string) string {
    return "[Unit]\n" +
        "Description=RustDesk Relay Server\n" +
        "After=network.target\n" +
        "[Service]\n" +
        "Type=simple\n" +
        "ExecStart=" + filepath.Join(installDir, "hbbr") + " -p " + hbrPort + "\n" +
        "WorkingDirectory=" + installDir + "\n" +
        "Restart=always\n" +
        "RestartSec=10\n" +
        "StandardOutput=append:" + filepath.Join(logDir, "relay.log") + "\n" +
        "StandardError=append:" + filepath.Join(logDir, "relay.error") + "\n" +
        "[Install]\n" +
        "WantedBy=multi-user.target\n"
}

func EnableAndStart() error {
	if err := util.RunCommandSilent("systemctl", "daemon-reload"); err != nil {
		return err
	}
	return util.RunCommandSilent("systemctl", "enable", "--now", "rustdesksignal.service", "rustdeskrelay.service")
}

func DisableAndStop() error {
	return util.RunCommandSilent("systemctl", "disable", "--now", "rustdesksignal.service", "rustdeskrelay.service")
}

func RemoveUnits(systemdDir string) error {
	_ = os.Remove(filepath.Join(systemdDir, "rustdesksignal.service"))
	_ = os.Remove(filepath.Join(systemdDir, "rustdeskrelay.service"))
	return util.RunCommandSilent("systemctl", "daemon-reload")
}
