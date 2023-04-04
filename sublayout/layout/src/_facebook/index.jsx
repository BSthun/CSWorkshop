import {
	Paper,
	Stack,
} from '@mui/material';
import React from 'react';

const Facebook = () => {
	return (
		<Stack>
			<Paper square sx={{ height: 64, zIndex: 20 }} elevation={2}>
				<iframe src="https://csc105-workshop.bsthun.com/components/comp01"
				        height="100%"
				        width="100%"
				        frameBorder="0"
				/>
			</Paper>
			<Stack flexDirection="row" justifyContent="space-between" height="calc(100vh - 64px)">
				<Paper square sx={{ width: 300 }} elevation={2}>
					<iframe src="https://csc105-workshop.bsthun.com/components/comp02"
					        height="100%"
					        width="100%"
					        frameBorder="0"
					/>
				</Paper>
				<Paper square sx={{ width: 600 }} elevation={0}>
					<iframe src="https://csc105-workshop.bsthun.com/components/comp03"
					        height="200px"
					        width="100%"
					        frameBorder="0"
					        style={{border: "1px solid #dadce0", marginTop: 8}}
					/>
					<iframe src="https://csc105-workshop.bsthun.com/components/comp04"
					        height="64px"
					        width="100%"
					        frameBorder="0"
					        style={{border: "1px solid #dadce0"}}
					/>
					<iframe src="https://csc105-workshop.bsthun.com/components/comp05"
					        height="600px"
					        width="100%"
					        frameBorder="0"
					        style={{border: "1px solid #dadce0"}}
					/>
				</Paper>
				<Paper square sx={{ width: 300 }} elevation={2}>
					<iframe src="https://csc105-workshop.bsthun.com/components/comp06"
					        height="120px"
					        width="100%"
					        frameBorder="0"
					/>
					<iframe src="https://csc105-workshop.bsthun.com/components/comp07"
					        height="250px"
					        width="100%"
					        frameBorder="0"
					/>
					<iframe src="https://csc105-workshop.bsthun.com/components/comp08"
					        height="500px"
					        width="100%"
					        frameBorder="0"
					/>
				</Paper>
			</Stack>
		</Stack>
	);
};

export default Facebook;
