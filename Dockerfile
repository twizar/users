FROM amazon/aws-lambda-go:1.2021.10.14.14

# Copy function code
COPY users_bin ${LAMBDA_TASK_ROOT}

# Set the CMD to your handler (could also be done as a parameter override outside of the Dockerfile)
CMD [ "users_bin" ]