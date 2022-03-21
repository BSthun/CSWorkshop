import {
	Container,
	Paper,
	Stack,
} from '@mui/material';
import React from 'react';

const Twitter = () => {
	return (
		<Container maxWidth="xl">
			<Stack flexDirection="row" height="100vh">
				<Paper square sx={{ flex: 5 }}>
					<Stack justifyContent="space-between" height="100%">
						<iframe src="https://csc105-workshop.bsthun.com/components/comp09"
						        frameBorder="0"
						        style={{ flex: 1 }}
						/>
						<iframe src="https://csc105-workshop.bsthun.com/components/comp10"
						        frameBorder="0"
						        height="120px"
						/>
					</Stack>
				</Paper>
				<Paper square sx={{ flex: 12 }}>
					<Stack justifyContent="space-between" height="100%">
						<iframe src="https://csc105-workshop.bsthun.com/components/comp11"
						        frameBorder="0"
						        height="200px"
						/>
						<iframe src="https://csc105-workshop.bsthun.com/components/comp12"
						        frameBorder="0"
						        height="500px"
						/>
						<iframe src="https://csc105-workshop.bsthun.com/components/comp14"
						        frameBorder="0"
						        height="200px"
						/>
					</Stack>
				</Paper>
				<Paper square sx={{ flex: 5 }}>
					<Stack>
						<iframe src="https://csc105-workshop.bsthun.com/components/comp16"
						        frameBorder="0"
						        height="500px"
						/>
						<iframe src="https://csc105-workshop.bsthun.com/components/comp17"
						        frameBorder="0"
						        height="300px"
						/>
					</Stack>
				</Paper>
			</Stack>
		</Container>
	);
};

export default Twitter;
