# Chapter 7: COMPUTER HARDWARE

1. `RAM`: Random Access Memory
2. `CPU`: CENTRAL PROCESSING UNIT
3. `Secondary Storage` : HDD / SSD
4. `(I/O)` : Input/ Output

## 1. RAM

- `SRAM`: Static Random Access Memory (a type of flip-flop which keep bites while power is connected)
- `DRAM`: Dynamic Random Access Memory (using a transistor and a capacitor)
- The capacitor’s charge `leaks` over time, so data must be `rewritten` to the cells.
- This refreshing of the memory cells is what makes DRAM `dynamic`.

#### SRAM vs DRAM:

- `DRAM` is low price so commonly used
- `SRAM` is faster but expensive so used for `cache memory` which speed is important.

#### physical memory addresses

- RAM is grids of memory cells and the location of each single bit cell in a grid can be identified using two-dimensional coordinates.
- RAM accesses multiple grids of 1-bit memory cells in parallel, allowing for reads or writes of multiple bits at once.
- Physical addresses stored the location of `8 bits or a byte of data` which is in binary format(certain number of bits) to identify each data is stored in memory cell in the RAM.
- These addresses allow the CPU to directly interact with the main memory(RAM) to read and write data.
- That number of bits in physical addresses determines not only the size of each memory address, but also the range of addresses available for the computer hardware to use the physical address space.

  ##### Computers with 64KB of memory:

  1. 64KB RAM = 64 \* (2 ^ 10) = 64 \* 1024= 65,536 B
  2. To find the range of memory addresses or how many bits need, we calculate `count combination`
  3. 2 ^ n = 65,536 ===> n = log2(65,536) = 16 (base-2 logarithm)
  4. 16 bits are required for unique memory addresses

  - From 0 (decimal) or 0000 (hex) or 0000 0000 0000 0000 (bits)
  - To 65,535 (decimal) or FFFF (hex) or 1111 1111 1111 1111 (bits)

## 2. CPU

#### Instruction Set Architectures (ISA)

- ISA is a model of how a CPU works.

1. `x86`: naming by Intel uses for computers, laptop, servers

   - Intel processors: 8086, 80186, 80286, 80386, 80486, Pentium and Celeron.
   - Software developed for an older x86 CPU runs on a newer x86 CPU, but software built for a newer x86 CPU that takes advantage of new x86 instructions won’t be able to run on older x86 CPUs that don’t understand the new instructions.
   - The x86 architecture has been 16-bit, 32-bit, and 64-bit.
   - A 32-bit CPU has 32-bit `registers`, a 32-bit `address bus`, or a 32-bit `data bus` and can operate on values that are 32 bits in length.

2. `ARM`: uses for smartphones and tablets because of reduced power consumption and lower cost as compared to x86.

   - To be used in system- on-chip (SoC ) designs, where a single integrated circuit contains not only a CPU, but also memory and other hardware.

#### CPU Components

1.  `Processor registers`:

    - To hold data during processing because accessing registers is very faster compared to accessing main memory.
    - Registers can only hold very small amounts of data in bits, not bytes because registers are so small.
    - The memory cells used in registers are a type of SRAM (flip-flop).
    - A 32-bit CPU can hold 32 bits of data.

2.  `Arithmetic Logic Unit`:

    - To performs logical and mathematical operations.

3.  `Control unit`:

    - To directs the CPU, communicating with the processor registers, the ALU, and main memory.

4.  `Clock`:

    - Each pulse of the clock is a signal to the CPU to move forward with executing a program.
    - Modern CPUs have clock speeds measured in gigahertz (GHz).
    - e.g. A 2GHz CPU has a clock that oscillates at 2 billion cycles per second

5.  `Cores`:

    - CPUs can perform more instructions per second by increasing the frequency of the clock but CPUs have a limit on their input clock frequency because beyond that limit leads to excessive heat generation or logic gates crashes.
    - Now new approach is focus on execution of multiple instructions in parallel rather than increasing clock frequency.
    - A CPU with multiple processing units called cores and each core is an independent processor that resides alongside other independent processors in a single CPU package.

6.  `Cache`:

    - Some instructions take multiple cycles to complete so CPUs tend to access the same memory locations over and over which is inefficient.
    - To avoid this inefficient, a small amount of memory (cache) resides within the CPU that holds a copy of data frequently accessed from main memory.
    - cache has multiple levels, often three. L1, L2 and L3.
    - L1 cache is the fastest to access, but it’s also the smallest.
    - L2 is slower and larger.
    - L3 is slower and larger still.
    - It is still slower to access main memory than any level of cache
    - Caches can be specific to each core or shared among the cores.

#### How CPU works:

`PC`: program counter is a register within the CPU that acts as an instruction pointer. It holds the `memory address` of the next instruction to be fetched and executed by the CPU.

1. Control unit `checks cache` from L1 then L2 then L3 for the needed data or instructions.
2. If data wasn't in the cache, control unit `finds the address` of an instruction in main memory(RAM) by PC (program counter).
3. `Fetch(read)` the instruction from `physical memory addresses`.
4. `Store` it in the `instruction register`.
5. `Update` the PC to point to the next instruction.
6. `decode` the current instruction for making sense of the 1 and 0 that represent.
7. `execute` the instruction which may require ALU to perform it (e.g. for addition operation)

#### pipelining in CPU:

- Allowing multiple instructions to be executed in parallel by a `single core`.
- pipelining only happens within `each core` to divide instructions into smaller steps.
- Multiple cores running in parallel is not the same as pipelining. It means that each core works on a different task or a separate set of instructions.
- adding multiple cores to a computer’s CPU doesn’t mean all applications benefit immediately or equally.
- `Software must be written to take advantage of parallel processing of instructions to get the maximum benefit of multi core hardware. Also modern operating systems run multiple programs at once.`

## 3. Secondary Storage (HDD / SSD)

- Memory(RAM) in a computer loosing data every time when device was powered down even OS or installed application.

- Computers have secondary storage that remembers data even when the system is powered down.

- Secondary storage is not directly addressable by the CPU unlike RAM.

- When a computer is powered on, the OS loads from secondary storage into main memory(RAM). The same for any applications that are set to run at startup or any user data (documents, music, settings) before it can be used.

- It is much cheaper per byte than RAM.

- Allowing for a large capacity of storage as compared to RAM.

- It is slower than RAM.

#### HDD vs SSD

- `hard disk drive (HDD)` stores data using magnetism on a rapidly spinning platter.
- `solid-state drive (SSD)` stores data using electrical charges in nonvolatile memory cells.

- SSDs are faster, quieter, and more resistant to mechanical failure compared to HDDs

a computer can load data on demand. When a computer is powered on, the operating system loads from secondary storage into main memory

## 4. Input/ Output (I/O)

- `Human interaction`:
  - To receive input from the outside (keyboard, mouse)
  - To send data to the outside (monitor, printer)
  - Both (touchscreen)
- `Computer-to-computer interaction`:
  - From a computer network such as the internet

#### Communicate with I/O devices

- Both way refer to a device controller rather than directly to data stored on the device.

1. `MMIO`:

   - The CPU requires to communicate with I/O devices so that devices need `physical memory addresses` to CPU can interact with them by reading or writing.
   - When physical address space is mapped to an I/O device and assigned address(es) is known as Memory-Mapped I/O (MMIO)
   - Addresses in physical address space don’t always refer to bytes of memory; they can also refer to an I/O device.

2. `PMIO`:

   - Some CPU families are assigned an I/O port to access I/O devices. A port is like a memory address, but instead of referring to a location in memory, the port number refers to an I/O device.
   - Accessing I/O devices through a separate port address space is known as port-mapped I/O (PMIO)

#### Bus

- To communicate between the CPU, memory, and I/O devices.
- A bus is a hardware communication system used by computer components.
- three common bus types:
  1. `Address Bus`: It acts as a selector for the memory address that the CPU wishes to access.
  2. `Data Bus`: It transmits a value read from memory or a value to be written to memory.
  3. `Control Bus`: It manages the operations happening over the other two buses. It can carry a signal indicating the status of an operation.
