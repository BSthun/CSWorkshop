import { Check } from '@mui/icons-material'
import { Box, Typography } from '@mui/material'
import React from 'react'

type TaskItemListProps = {
	name: string
	finished: boolean
	selectedIndex: number
	index: number
}

const TaskItemList: React.FC<TaskItemListProps> = ({
	name,
	finished,
	selectedIndex,
	index,
}) => {
	return (
		<Box
			sx={{
				display: 'flex',
				justifyContent: 'space-between',
				alignItems: 'center',
			}}
		>
			<Typography
				sx={{
					color: selectedIndex === index ? 'black' : '#7E7E7E',
					pr: 2,
				}}
			>
				{name}
			</Typography>
			{finished && <Check sx={{ color: '#29D162' }} />}
		</Box>
	)
}

export default TaskItemList
