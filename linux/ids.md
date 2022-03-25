## 关于 ID

于编程中，通常需要使用唯一标识符，用于绑定到各种硬件和软件对象以及软件生命周期。通常，当人们查找要使用的此类 ID 时，他们会选择错误的 ID，因为语义和生命周期或 ID 不明确。下面是一些在 Linux 上可访问的 ID 以及如何使用它们的简略指南。

## 硬件 ID

/sys/class/dmi/id/product_uuid：  
主板product UUID，由主板制造商设置并编码在 BIOS DMI 信息中。它可以用于识别主板并且仅用于识别主板。当用户更换主板时它会改变。此外，经常有 BIOS 制造商在其中写入伪造的序列号。此外，它是特定于 x86 架构的。禁止非特权用户访问。  
位于/sys/class/dmi/id/product_uuid目录，同级目录下还有其他一些关于主板硬件相关的信息。
dmidecode命令同样可以获得相关的id信息，如下：
```bash
$ dmidecode -s
dmidecode: option requires an argument -- 's'
String keyword expected
Valid string keywords are:
  bios-vendor
  bios-version
  bios-release-date
  system-manufacturer
  system-product-name
  system-version
  system-serial-number
  system-uuid
  system-family
  baseboard-manufacturer
  baseboard-product-name
  baseboard-version
  baseboard-serial-number
  baseboard-asset-tag
  chassis-manufacturer
  chassis-type
  chassis-version
  chassis-serial-number
  chassis-asset-tag
  processor-family
  processor-manufacturer
  processor-version
  processor-frequency

$ sudo dmidecode -s system-uuid
ddd20f6d-8484-2384-ff6a-04d9f51b8628
```

CPUID/EAX=3 CPU serial number：  
CPU UUID，由 CPU 制造商设置并在 CPU 芯片上编码。它可以用来识别一个 CPU 并且只识别一个 CPU。当用户更换 CPU 时它会发生变化。此外，大多数现代 CPU 不再实现此功能，并且较旧的计算机默认情况下倾向于禁用此选项，可通过 BIOS 设置选项进行控制。此外，它是特定于 x86 的。
dmidecocde -t process可以查看。

/sys/class/net/*/address：  
一个或多个网络 MAC 地址，由网络适配器制造商设置并在某些网卡 EEPROM 上编码。当用户更换网卡时它会改变。由于网卡是可选的，如果不能保证此 ID 的可用性可能不止一张，而且您可能有不止一张可供选择。在虚拟机上，MAC 地址往往是随机的。因此，这也没有什么普遍用途。

/sys/bus/usb/devices/*/serial：  
各种 USB 设备的序列号，编码在 USB 设备 EEPROM 中。大多数设备没有设置序列号，如果有，通常是伪造的。如果用户更换他的 USB 硬件或将其插入另一台机器，这些 ID 可能会更改或出现在其他机器中。因此，这也没什么用。
还有各种其他可用的硬件 ID，您可以通过各种设备（如硬盘和类似设备）的 ID_SERIAL udev 属性发现其中的许多。它们都有一个共同点，即它们都绑定到特定的（可替换的）硬件，不是普遍可用的，通常充满虚假数据并且在虚拟化环境中是随机的。或者换句话说：不要使用它们，不要依赖它们来识别，除非你真的知道你在做什么，而且一般来说它们不能保证你可能希望他们保证的东西。

## 软件 ID

/proc/sys/kernel/random/boot_id：  
在每次启动时重新生成的随机 ID。因此，它可以用来识别本地机器的当前引导。它在任何最新的 Linux 内核上普遍可用。如果您需要识别特定引导内核上的特定引导，这是一个不错且安全的选择。

gethostname() , /proc/sys/kernel/hostname：  
管理员配置的非随机 ID，用于识别网络中的机器。通常这根本没有设置或设置为一些默认值，例如 localhost甚至在本地网络中都不是唯一的。此外，它可能会在运行时发生变化，例如因为它会根据更新的 DHCP 信息而变化。因此，除了向用户展示之外，它几乎完全没有用处。它的语义很弱，并且依赖于管理员的正确配置。不要使用它来识别分布式环境中的机器。除非集中管理，否则它不会起作用，这使得它在全球化的移动世界中毫无用处。它在应绑定到特定主机的自动生成的文件名中没有位置。请不要使用它。真的不是很多人想的那样。 gethostname()在 POSIX 中是标准化的，因此可以移植到其他 Unix。

IP 地址：这些往往是动态分配的，并且通常仅在本地网络上有效，甚至仅在本地链接上有效（即 192.168.xx 样式地址，甚至 169.254.xx/IPv4LL）。不幸的是，它们因此在网络之外几乎没有用处。

gethostid()：  
返回当前机器的假定唯一的 32 位标识符。其语义尚不清楚。在大多数机器上，这只是返回一个基于本地 IPv4 地址的值。在其他情况下，它是通过/etc/hostid文件由管理员控制的。由于此 ID 的语义不明确，而且通常只是基于 IP 地址的值，因此使用它几乎总是错误的选择。最重要的是32位并不是特别多。另一方面，这在 POSIX 中是标准化的，因此可以移植到其他 Unix。最好忽略这个值，如果人们不想忽略它，他们可能应该将/etc/hostid符号链接到 /var/lib/dbus/machine-id或类似的东西。

/var/lib/dbus/machine-id：  
标识特定 Linux/Unix 安装的 ID。如果更换硬件，它不会改变。它在虚拟化环境中并非不可靠。该值具有清晰的语义，被认为是 D-Bus API 的一部分。它被认为是全球唯一的，并且可移植到所有具有 D-Bus 的系统。在 Linux 上，它是普遍可用的，因为现在几乎所有非嵌入式甚至相当一部分嵌入式机器都提供 D-Bus。这是识别机器的推荐方法，可能会回退到主机名以覆盖仍然缺乏 D-Bus 的系统。如果您的应用程序链接到libdbus ，您可以使用dbus_get_local_machine_id()访问此 ID ，否则您可以直接从文件系统中读取它。
cat /var/lib/dbus/machine-id

/proc/self/sessionid：  
标识特定 Linux 登录会话的 ID。此 ID 由内核维护，也是审计逻辑的一部分。它在特定系统启动期间唯一分配给每个登录会话，由会话的每个进程共享，即使跨 su/sudo 也是如此，并且不能由用户空间更改。不幸的是，到目前为止，某些发行版未能正确设置它以使其正常工作（嘿，你，Ubuntu！），并且此 ID 始终为 (uint32_t) -1。但希望他们最终能解决这个问题。尽管如此，对于本地机器上的唯一会话标识符和当前引导来说，它是一个不错的选择。要使此 ID 全局唯一，最好与/proc/sys/kernel/random/boot_id结合使用。

getuid()：  
标识特定 Unix/Linux 用户的 ID。此 ID 通常在创建用户时自动分配。它在机器之间不是唯一的，如果原始用户被删除，它可能会重新分配给不同的用户。因此，它应该只在本地使用，并考虑到有限的有效性。要使此 ID 全局唯一，将其与/var/lib/dbus/machine-id组合是不够的，因为相同的 ID 可能用于稍后使用相同 UID 创建的不同用户。尽管如此，这种组合通常已经足够好了。它适用于所有 POSIX 系统。id命令可查看。

ID_FS_UUID：  
标识 udev 树中特定文件系统的 ID。这些序列是如何生成的并不总是很清楚，但这往往在几乎所有现代磁盘文件系统上都可用。它不适用于 NFS 挂载或虚拟文件系统。尽管如此，这通常是识别文件系统的好方法，在根目录的情况下甚至是安装。然而，由于定义较弱的生成语义，通常首选 D-Bus 机器 ID。


## Generating IDs
Linux 提供了一个内核接口，通过读取/proc/sys/kernel/random/uuid来按需生成 UUID 。这是一个非常简单的生成 UUID 的接口。也就是说，UUID 背后的逻辑过于复杂，通常使用 uuidgen读取 16 个字节左右是更好的选择。
cat /proc/sys/kernel/random/uuid

## 概括
使用/var/lib/dbus/machine-id！
使用 /proc/self/sessionid！
使用/proc/sys/kernel/random/boot_id！
使用getuid()！使用/dev/urandom！
忘记其余的，尤其是主机名或硬件 ID，例如 DMI。请记住，您可以通过各种方式组合上述 ID，以获得不同的语义和有效性约束。

## 参考
http://0pointer.de/blog/projects/ids.html
https://nyogjtrc.github.io/posts/2018/12/some-unique-id-in-linux/