# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from .. import _utilities

concurrency: int
"""
Maximum number of hosts to configure in parallel, set to 0 for unlimited
"""

concurrentUploads: int
"""
Maximum number of files to upload in parallel, set to 0 for unlimited
"""

noDrain: bool
"""
Do not drain worker nodes when upgrading
"""

noWait: bool
"""
Do not wait for worker nodes to join
"""

skipDowngradeCheck: bool
"""
Skip downgrade check
"""

