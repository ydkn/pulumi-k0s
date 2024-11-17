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

import types

__config__ = pulumi.Config('k0s')


class _ExportableConfig(types.ModuleType):
    @property
    def concurrency(self) -> int:
        """
        Maximum number of hosts to configure in parallel, set to 0 for unlimited
        """
        return __config__.get_int('concurrency') or (_utilities.get_env_int('PULUMI_K0S_CONCURRENCY') or 30)

    @property
    def concurrent_uploads(self) -> int:
        """
        Maximum number of files to upload in parallel, set to 0 for unlimited
        """
        return __config__.get_int('concurrentUploads') or (_utilities.get_env_int('PULUMI_K0S_CONCURRENT_UPLOADS') or 5)

    @property
    def no_drain(self) -> bool:
        """
        Do not drain worker nodes when upgrading
        """
        return __config__.get_bool('noDrain') or (_utilities.get_env_bool('PULUMI_K0S_NO_DRAIN') or False)

    @property
    def no_wait(self) -> bool:
        """
        Do not wait for worker nodes to join
        """
        return __config__.get_bool('noWait') or (_utilities.get_env_bool('PULUMI_K0S_NO_WAIT') or False)

    @property
    def skip_downgrade_check(self) -> bool:
        """
        Skip downgrade check
        """
        return __config__.get_bool('skipDowngradeCheck') or (_utilities.get_env_bool('PULUMI_K0S_SKIP_DOWNGRADE_CHECK') or False)

