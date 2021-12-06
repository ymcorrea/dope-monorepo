
pragma solidity ^0.8.0;

contract GetMenHair {
    function getRles() external returns (uint256, bytes[] memory) {
        bytes[] memory rles = new bytes[](19);
        bytes[19] memory _rles = [bytes(hex''),bytes(hex'0009230f1f0100012f0247012f0300012f030004000400012f0300'),bytes(hex'0009230e1f0200024d014d0100024d014d03000400014d0300'),bytes(hex'0009230e1f0100014a024b014b0100024b014b03000400014a0300'),bytes(hex'0009250c1e0200012f024e02000100012f0100024e0100014e022f0400014e'),bytes(hex'0009250c1e0200014502460200010001460100024601000146024604000146'),bytes(hex'0008230b1f0200012f014c0100022f014c012f0200014c'),bytes(hex'0008250c1f0100022f014f012f0100032f024f012f012f0300014f012f012f0500'),bytes(hex'0009250e1f01000148034901000248020001490100014804000149060001480500'),bytes(hex'0007250e1e0200012c020302000100022c03030100032c0403022c04000103022c04000103012c06000100012c0500'),bytes(hex'0006260f1d0200022c030302000100032c04030100042c0503042c0503032c04000203032c040001030100022c07000100022c06000200012c0600'),bytes(hex'0009240e1f0100012f022e0100012f0300012e012f04000500012f0400'),bytes(hex'0007240c1f0400012f0200032f052f0500012f0400'),bytes(hex'000926111d0300012d032f02000200012d0400012f01000200012d0400012f01000100012d0600012f012d0100012d0500012f0200012d06000100012d0500012f01000100012d0600012f'),bytes(hex'000923101f0100012f0100012f0400012f03000400012f0300012f0300012f0300'),bytes(hex'0007240e1d0100032f0300042f0300062f01000100022f0300012f0200012f040007000200012f0400'),bytes(hex'000825111f0100012f032e0100022f042e012f0100012e0100022e012f0400012e06000600012f0500012f0500012f0500'),bytes(hex'000825111f0100012f032e0100032f032e0300012f022e060006000600012f0500012f0500012f0500'),bytes(hex'000825111f010001500351010003500351030001500251060006000600015205000152050001520500')];
        for (uint256 i = 0; i < rles.length; i++) {
            rles[i] = _rles[i];
        }

        return (2, rles);
    }
}