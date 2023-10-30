namespace DS_Lab1.CommunicationElements
{
    public class ErrorResponse
    {
        public string? Message { get; set; }

        public ErrorResponse(string? message)
        {
            Message = message;
        }
    }
}
