import { AccessTime, Storage } from '@mui/icons-material'
import { Box, Typography } from '@mui/material'

type EnrollmentProp = {
	labName: string
	dbName: string
	enrollAt: Date
	dbValid: boolean
}

const EnrollmentBox: React.FC<EnrollmentProp> = (props) => {
	const dateFormat = (date: Date): string => {
		var options = {
			weekday: 'long',
			year: 'numeric',
			month: 'long',
			day: 'numeric',
		}

		return date.toLocaleDateString('th-TH')
	}
	return (
		<Box
			sx={{
				width: '300px',
				height: '180px',
				boxShadow: 'rgba(0, 0, 0, 0.24) 0px 3px 8px;',
				borderRadius: '12px',
				padding: '20px',
				'& > *': {
					marginY: '10px',
				},
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
			<Box>
				<Typography fontSize={22}>{props.labName}</Typography>
			</Box>
			<Box
				sx={{
					display: 'flex',
					alignItems: 'center',
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
					{dateFormat(props.enrollAt) +
						' ' +
						props.enrollAt.toLocaleTimeString()}
				</Typography>
			</Box>
		</Box>
	)
}

export default EnrollmentBox
