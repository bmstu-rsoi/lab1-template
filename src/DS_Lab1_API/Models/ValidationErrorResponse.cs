namespace DS_Lab1.CommunicationElements
{
    public class ValidationErrorResponse
    {
        public string? Message { get; set; }
        public Dictionary<string, string>? Errors { get; set; }
    }
}
