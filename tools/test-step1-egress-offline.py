#!/usr/bin/env python3
"""Regression: explicit offline egress mode must never spawn curl."""

import importlib.util
from importlib.machinery import SourceFileLoader
import inspect
from pathlib import Path
import unittest
from unittest.mock import patch


SCRIPT = Path(__file__).with_name("adint-step1-run")
LOADER = SourceFileLoader("adint_step1_run", str(SCRIPT))
SPEC = importlib.util.spec_from_loader("adint_step1_run", LOADER)
MODULE = importlib.util.module_from_spec(SPEC)
SPEC.loader.exec_module(MODULE)


class OfflineEgressTest(unittest.TestCase):
    def test_live_egress_probe_remains_the_default(self):
        parameters = inspect.signature(MODULE.egress_record).parameters
        self.assertIn("probe_external", parameters)
        if "probe_external" in parameters:
            self.assertIs(parameters["probe_external"].default, True)

    def test_offline_record_is_unknown_and_never_runs_external_probe(self):
        parameters = inspect.signature(MODULE.egress_record).parameters
        self.assertIn("probe_external", parameters,
                      "egress_record needs an explicit offline mode")
        if "probe_external" not in parameters:
            return

        with patch.object(MODULE.subprocess, "run",
                          side_effect=AssertionError("offline mode invoked subprocess")) as run:
            record = MODULE.egress_record(probe_external=False)

        run.assert_not_called()
        self.assertEqual(record["probe_status"], "unknown")
        self.assertEqual(record["probe_reason"], "disabled by explicit offline mode")
        self.assertIsNone(record["legs_agree"])
        self.assertNotIn("ip", record["node_egress"])
        self.assertNotIn("ip", record["node_egress_direct"])


if __name__ == "__main__":
    unittest.main()
