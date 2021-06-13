package msmnile

import (
    "android/soong/android"
)

func init() {
    android.RegisterModuleType("oneplus_msmnile_bootctrl_library", bootctrlLibraryFactory)
}
