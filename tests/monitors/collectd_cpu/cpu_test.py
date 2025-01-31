"""
Tests for the collectd/cpu monitor
"""
import time

import pytest

from tests.helpers.agent import Agent
from tests.helpers.metadata import Metadata
from tests.helpers.verify import run_agent_verify_all_metrics

pytestmark = [pytest.mark.collectd, pytest.mark.cpu, pytest.mark.monitor_without_endpoints]

METDATA = Metadata.from_package("collectd/cpu")


def test_collectd_cpu_default():
    with Agent.run(
        """
        monitors:
        - type: collectd/cpu
        """
    ) as agent:
        # There aren't any default.
        time.sleep(15)
        assert not agent.fake_services.datapoints


def test_collectd_cpu_all():
    run_agent_verify_all_metrics(
        """
        monitors:
        - type: collectd/cpu
          extraMetrics: ["*"]
        """,
        METDATA,
    )
