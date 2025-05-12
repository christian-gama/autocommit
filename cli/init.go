package cli

func init() {
	AutoCommit.AddCommand(Configure)
	AutoCommit.AddCommand(Instruction)
	Instruction.AddCommand(restoreInstruction)
}
