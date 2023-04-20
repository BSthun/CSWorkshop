import {
	Autorenew,
	BatchPrediction,
	FormatListBulleted,
} from '@mui/icons-material'
import {
	Box,
	Button,
	CircularProgress,
	Dialog,
	DialogContent,
	DialogTitle,
	Divider,
	IconButton,
	Stack,
	Table,
	TableBody,
	TableCell,
	TableContainer,
	TableHead,
	TableRow,
	TextField,
	Typography,
} from '@mui/material'
import React, { useCallback, useRef, useState } from 'react'
import TaskItemList from '../components/TaskItemList'
import axios from 'axios'
import { BasedResponse } from '../types/APIs/basedResponse'
import { EnrollmentInfoAPI } from '../types/APIs/Lab/enrollment_info'
import { LabState, WebsocketMessage } from '../types/APIs/Lab/lab_state'
import { useParams } from 'react-router-dom'
import MockComponent from '../components/MockComponent'
import { toast } from 'react-toastify'

interface DbInfo {
	label: string
	value: string
}

const taskTags = [
	{
		name: 'Level',
		value: 'Lv1',
	},
	{
		name: 'Difficulty',
		value: 'Easy',
	},
	{
		name: 'Tag 3',
		value: 'Hello',
	},
]

const Lab = () => {
	const [termKey, setKey] = useState(String(Math.random()))
	const [isTask, setIsTask] = React.useState(false)
	const [selectedTask, setSelectedTask] = React.useState(0)
	const [isLoading, setIsLoading] = React.useState(false)
	const [labState, setLabState] = React.useState<LabState | null>(null)
	const [enrollmentInfo, setEnrollmentInfo] =
		React.useState<EnrollmentInfoAPI>()
	const [dbInfo, setDbInfo] = React.useState<DbInfo[]>([])
	const params = useParams()
	const [hint, setHint] = useState<string | null>(null)
	const websocketRef = useRef<WebSocket>()

	const fetchLabInfo = async () => {
		try {
			const infoData = await axios.get<BasedResponse<EnrollmentInfoAPI>>(
				`/api/lab/enroll/info?enrollmentId=${params.enrollmentId}`
			)
			setEnrollmentInfo(infoData.data.data)
			setDbInfo([
				{
					label: 'Database Name',
					value: infoData.data.data.dbName ?? 'error',
				},
				{
					label: 'Database Host',
					value: infoData.data.data.dbHost ?? 'error',
				},
				{
					label: 'Database Port',
					value: infoData.data.data.dbPort ?? 'error',
				},
				{
					label: 'Database Username',
					value: infoData.data.data.dbUsername ?? 'error',
				},
				{
					label: 'Database Password',
					value: infoData.data.data.dbPassword ?? 'error',
				},
			])
			initializeWebsocket(
				infoData.data.data.enrollmentId,
				infoData.data.data.token
			)
		} catch {
			toast.error(e.response?.data?.message ?? 'Fetch lab info error')
		}
	}

	const clickTask = async (taskId: number) => {
		try {
			await axios.get(
				`/api/lab/task/click?enrollmentId=${params.enrollmentId}&taskId=${taskId}`
			)
		} catch (e) {
			toast.error(e.response?.data?.message ?? 'Click task error')
		}
	}

	const initializeWebsocket = (enrollmentId: number, token: string) => {
		let timer: any
		if (websocketRef.current) {
			websocketRef.current?.close()
		}
		websocketRef.current = new WebSocket(
			`${
				import.meta.env.VITE_BACKEND_WEBSOCKET_URL
			}/ws/lab?eid=${enrollmentId}&token=${token}`
		)
		websocketRef.current.onopen = (e) => {
			console.log('open')
			timer = setInterval(() => {
				websocketRef.current?.send(
					JSON.stringify({
						type: 'ping',
					})
				)
			}, 10 * 1000)
		}

		websocketRef.current.onmessage = (e: MessageEvent) => {
			const data = JSON.parse(e.data) as WebsocketMessage<LabState>
			setLabState(data.payload)
		}
		websocketRef.current.onclose = (e) => {
			clearInterval(timer)
		}
	}

	const showHint = () => {
		axios
			.get(
				`/api/lab/hint/text?enrollmentId=${
					params.enrollmentId
				}&taskId=${selectedTask + 1}`
			)
			.then((res) => {
				setHint(res.data.data.hint_text)
			})
			.catch((e) => {
				toast.error(e.response?.data?.message ?? 'Click task error')
			})
	}
	React.useEffect(() => {
		fetchLabInfo()
		return websocketRef.current?.close()
	}, [])

	return (
		<Box
			sx={{
				display: 'flex',
				height: '100%',
			}}
		>
			<Box
				sx={{
					width: 400,
					backgroundColor: 'white',
					boxShadow: '0px 0px 12px rgba(0, 0, 0, 0.25)',
					py: 4,
					overflowY: 'scroll',
				}}
			>
				<Box
					sx={{
						display: 'flex',
						alignItems: 'center',
						justifyContent: 'space-between',
						mb: 3.5,
						px: 4,
					}}
				>
					<Typography fontSize={24}>Spotify Test</Typography>
					<IconButton
						onClick={() => {
							setIsTask(!isTask)
						}}
					>
						<FormatListBulleted
							sx={{
								color: '#919191',
							}}
						/>
					</IconButton>
				</Box>
				{!isTask ? (
					<Box px={4}>
						{dbInfo.map((item, index) => (
							<TextField
								key={index}
								variant="outlined"
								label={item.label}
								fullWidth
								value={item.value}
								onChange={() => {}}
								sx={{
									my: 1,
								}}
							/>
						))}
					</Box>
				) : (
					<>
						{enrollmentInfo?.tasks.map((item, index) => (
							<Box key={index}>
								{index === 0 && <Divider />}
								<Box
									sx={{
										px: 4,
										py: 2.2,
										cursor: 'pointer',
										backgroundColor:
											selectedTask === index
												? '#F5F5F5'
												: 'white',
									}}
									onClick={async () => {
										setSelectedTask(index)
										await clickTask(item.id)
									}}
								>
									<TaskItemList
										name={item.title}
										finished={item.passed}
										index={index}
										selectedIndex={selectedTask}
									/>
								</Box>
								<Divider />
							</Box>
						))}
					</>
				)}
			</Box>
			<Box
				sx={{
					width: 'calc(100% - 400px)',
				}}
			>
				{labState == null ? (
					<Stack
						alignItems="center"
						justifyContent="center"
						height="100%"
					>
						<CircularProgress />
					</Stack>
				) : !labState.dbValid ? (
					<Box
						sx={{
							display: 'flex',
							justifyContent: 'center',
							alignItems: 'center',
							height: '100%',
							'& > *': {
								mx: 3,
							},
						}}
					>
						<MockComponent key={termKey} />
						<Box
							sx={{
								width: '180px',
								height: '180px',
								borderRadius: 2,
								backgroundColor: 'white',
								boxShadow: '0px 4px 4px rgba(0, 0, 0, 0.25)',
								display: 'flex',
								flexDirection: 'column',
								justifyContent: 'center',
								alignItems: 'center',
							}}
							onClick={() => {
								!isLoading && setIsLoading(true)
							}}
						>
							<Box
								sx={{
									width: '70px',
									height: '90px',
									display: 'flex',
									justifyContent: 'center',
									alignItems: 'center',
								}}
							>
								{isLoading ? (
									<CircularProgress color="info" />
								) : (
									<>
										<Autorenew
											sx={{
												width: '100%',
												height: '100%',
												color: '#919191',
											}}
										/>
									</>
								)}
							</Box>
							<Typography pt={1}>Recheck data</Typography>
						</Box>
					</Box>
				) : !labState.taskTitle ? (
					<Stack
						height="100%"
						alignItems="center"
						justifyContent="center"
					>
						Select a task to get started
					</Stack>
				) : (
					<Box
						sx={{
							height: 'calc(100vh - 64px)',
							overflowY: 'scroll',
						}}
					>
						<Box p={4}>
							<Typography fontSize={24} fontWeight={500} mb={1.5}>
								{labState.taskTitle}
							</Typography>
							<Typography fontSize={18} mb={1.5}>
								{labState.taskDescription}
							</Typography>
							<Stack direction="row" mb={4}>
								<Box flex={1}>
									{labState.taskTags.map((item, index) => (
										<Box display="flex" key={index}>
											<Typography
												fontFamily="monospace"
												mr={1.5}
												textTransform="uppercase"
											>
												{item.key}
											</Typography>
											<Typography
												fontFamily="monospace"
												fontWeight={700}
												textTransform="uppercase"
											>
												{item.value}
											</Typography>
										</Box>
									))}
								</Box>
								<Button
									variant="outlined"
									sx={{ alignSelf: 'flex-end' }}
									onClick={() => {
										showHint()
									}}
								>
									Show hint
								</Button>
							</Stack>
							<TextField
								variant="outlined"
								fullWidth
								multiline
								rows={4}
								value={labState.query || ''}
								onFocus={(e) => {
									e.preventDefault()
								}}
								sx={{
									'.MuiOutlinedInput-root': {
										borderRadius: 3,
										'& fieldset': {
											borderWidth: 3,
											borderColor: labState.queryPassed
												? '#4ABC4F'
												: '#BC4A4A',
										},
										'&:hover fieldset': {
											borderWidth: 3,
											borderColor: labState.queryPassed
												? '#4ABC4F'
												: '#BC4A4A',
										},
									},
									'.Mui-focused fieldset': {
										borderColor: labState.queryPassed
											? '#4ABC4F'
											: '#BC4A4A',
									},
									mb: 3,
								}}
							/>
							{labState?.queryPassed && (
								<Box
									sx={{
										backgroundColor: '#4ABC4F20',
										borderRadius: 3,
										p: 3,
										mb: 3,
									}}
								>
									<Typography
										fontFamily={'monospace'}
										sx={{
											color: '#3AAC3F',
										}}
									>
										Passed
									</Typography>
								</Box>
							)}
							{labState?.queryError && (
								<Box
									sx={{
										backgroundColor:
											'rgba(255, 118, 118, 0.2)',
										borderRadius: 3,
										p: 3,
										mb: 3,
									}}
								>
									<Typography fontFamily={'monospace'}>
										{labState?.queryError}
									</Typography>
								</Box>
							)}
							{labState?.result && (
								<Box
									sx={{
										display: 'flex',
									}}
								>
									<Box
										sx={{
											width: '50%',
										}}
									>
										<Typography mb={3}>
											Expected result
										</Typography>
										<TableContainer
											sx={{
												width: 'calc(100% - 16px)',
												border: '1px solid #eaeaea',
												borderRadius: 4,
											}}
										>
											<Table sx={{}}>
												<TableHead>
													<TableRow>
														{labState?.result.expectedHeader.map(
															(item, index) => (
																<TableCell
																	key={index}
																>
																	{item}
																</TableCell>
															)
														)}
													</TableRow>
												</TableHead>
												<TableBody>
													{labState?.result.expectedRows.map(
														(item, index) => (
															<TableRow
																key={index}
															>
																{item.map(
																	(
																		item,
																		index
																	) => (
																		<TableCell
																			key={
																				index
																			}
																		>
																			{
																				item
																			}
																		</TableCell>
																	)
																)}
															</TableRow>
														)
													)}
												</TableBody>
											</Table>
										</TableContainer>
									</Box>
									<Box
										sx={{
											width: '50%',
										}}
									>
										<Typography mb={3}>
											Actual result
										</Typography>
										<TableContainer
											sx={{
												width: 'calc(100% - 16px)',
												border: '1px solid #eaeaea',
												borderRadius: 4,
											}}
										>
											<Table>
												<TableHead>
													<TableRow>
														{labState?.result.actualHeader.map(
															(item, index) => (
																<TableCell
																	key={index}
																>
																	{item}
																</TableCell>
															)
														)}
													</TableRow>
												</TableHead>
												<TableBody>
													{labState?.result.actualRows.map(
														(item, index) => (
															<TableRow
																key={index}
															>
																{item.map(
																	(
																		item,
																		index
																	) => (
																		<TableCell
																			key={
																				index
																			}
																			sx={{
																				minWidth: 100,
																			}}
																		>
																			{
																				item
																			}
																		</TableCell>
																	)
																)}
															</TableRow>
														)
													)}
												</TableBody>
											</Table>
										</TableContainer>
									</Box>
								</Box>
							)}
						</Box>
					</Box>
				)}
			</Box>
			<Dialog onClose={() => setHint(null)} open={Boolean(hint)}>
				<DialogTitle>Hint</DialogTitle>
				<DialogContent>
					<Typography
						component="p"
						dangerouslySetInnerHTML={{
							__html: hint?.replace('\n', '<br>')!,
						}}
					/>
				</DialogContent>
			</Dialog>
		</Box>
	)
}

export default Lab
