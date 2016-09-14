from hadoopbeat import BaseTest

import os


class Test(BaseTest):

    def test_base(self):
        """
        Basic test with exiting Hadoopbeat normally
        """
        self.render_config_template(
                path=os.path.abspath(self.working_dir) + "/log/*"
        )

        hadoopbeat_proc = self.start_beat()
        self.wait_until( lambda: self.log_contains("hadoopbeat is running"))
        exit_code = hadoopbeat_proc.kill_and_wait()
        assert exit_code == 0
