# Calculate password for Hyundai/Kia Mobis updates
## Required params
**ro.product.model**=daudioplus,daudio, etc

**ro.product.brand**=hyundai or kia

**ro.product.name**=pde_eu ,tlfl_eu, sk_ru, etc

**ro.product.device**=${ro.product.model}low_${ro.product.name}
**ro.product.board**=daudio

**ro.product.cpu.abi**=armeabi-v7a

**ro.product.cpu.abi2**=armeabi

**ro.product.manufacturer**=mobis

**ro.product.locale.language**=en, ru, etc

**ro.product.locale.region**=GB, RU, US, etc

## Usage
   * UpdatePasswordCalculator.exe *pathToBuildProp*
   * Put build.prop file near UpdatePasswordCalculator.exe
