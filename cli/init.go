package cli

func init() {
	AutoCommit.AddCommand(Configure)
	Instruction.AddCommand(restoreInstruction)
	AutoCommit.AddCommand(Instruction)
}
