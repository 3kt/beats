import metricbeat
import os
import platform
import sys
import unittest


GOLANG_FIELDS = metricbeat.COMMON_FIELDS + ["golang"]


class Test(metricbeat.BaseTest):

    @unittest.skipIf(platform.platform().startswith("Windows-10"),
                     "flakey test: https://github.com/elastic/beats/issues/25858")
    def test_stats(self):
        """
        golang heap test
        """
        self.render_config_template(modules=[{
            "name": "golang",
            "metricsets": ["heap"],
            "hosts": ["http://localhost:6060"],
            "period": "1s"
        }])
        proc = self.start_beat(
            extra_args=["--httpprof", "localhost:6060"])

        self.wait_until(lambda: self.output_lines() > 0)
        proc.check_kill_and_wait()
        self.assert_no_logged_warnings()

        output = self.read_output_json()
        evt = output[0]

        self.assertCountEqual(self.de_dot(GOLANG_FIELDS), evt.keys(), evt)
        assert evt["golang"]["heap"]["allocations"]["total"] > 0

        self.assert_fields_are_documented(evt)
