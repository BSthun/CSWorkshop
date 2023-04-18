import { AccessTime, Storage } from '@mui/icons-material'
import { Box, Typography } from '@mui/material'
import { format } from 'date-fns'

type EnrollmentProp = {
	labName: string
	dbName: string
	enrollAt: Date
	dbValid: boolean
}

const EnrollmentBox: React.FC<EnrollmentProp> = (props) => {
	return (
		<Box
			sx={{
				width: '300px',
				height: '180px',
				boxShadow: '0px 2px 12px rgba(0, 0, 0, 0.25)',
				borderRadius: '12px',
				padding: '20px',
				backgroundColor: 'white',
			}}
		>
			<Box
				sx={{
					height: '30px',
					display: 'flex',
					alignItems: 'center',
				}}
			>
				<Box
					sx={{
						height: '3px',
						width: '200px',
						background: props.dbValid
							? 'rgba(139, 211, 128, 1)'
							: 'red',
						borderRadius: '10px',
					}}
				/>
			</Box>
			<Box mb={1.5}>
				<Typography fontSize={22}>{props.labName}</Typography>
			</Box>
			<Box
				sx={{
					display: 'flex',
					alignItems: 'center',
					mb: 1,
				}}
			>
				<Storage
					sx={{
						color: 'rgba(145, 145, 145, 1)',
						marginRight: '10px',
					}}
				/>
				<Typography fontSize={18} color={'rgba(145, 145, 145, 1)'}>
					{props.dbName}
				</Typography>
			</Box>
			<Box
				sx={{
					display: 'flex',
					alignItems: 'center',
				}}
			>
				<AccessTime
					sx={{
						color: 'rgba(145, 145, 145, 1)',
						marginRight: '10px',
					}}
				/>
				<Typography fontSize={18} color={'rgba(145, 145, 145, 1)'}>
					{format(new Date(props.enrollAt), 'MMMM dd yyyy, HH:mm a')}
				</Typography>
			</Box>
		</Box>
	)
}

export default EnrollmentBox
