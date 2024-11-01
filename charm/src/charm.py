#!/usr/bin/env python3
# Copyright 2024 Ubuntu
# See LICENSE file for licensing details.

"""Charm the application."""

import logging

import ops

logger = logging.getLogger(__name__)

ARTIFACT_URL = ""


class DarkAntichatMachineCharm(ops.CharmBase):
    """Charm the application."""

    def __init__(self, framework: ops.Framework):
        super().__init__(framework)
        self.framework.observe(self.on.start, self._on_start)
        self.framework.observe(self.on.install, self._on_install)

    def _on_install(self, event: ops.InstallEvent):
        """Handle install event."""
        self.unit.status = ops.MaintenanceStatus("Installing microsample snap")
        os.system(f"wget -O /usr/bin/dark-chat {ARTIFACT_URL}")
        os.system('chmod +x dark-chat')
        self.unit.status = ops.ActiveStatus("Ready")

    def _on_start(self, event: ops.StartEvent):
        """Handle start event."""
      os.system("/usr/bin/dark-chat &")

      # Everything is awesome
      self.unit.status = ops.ActiveStatus()


if __name__ == "__main__":  # pragma: nocover
    ops.main(DarkAntichatMachineCharm)  # type: ignore
